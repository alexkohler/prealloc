package pkg

import (
	"fmt"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

type sliceDeclaration struct {
	name string
	pos  token.Pos
}

type returnsVisitor struct {
	// flags
	simple            bool
	includeRangeLoops bool
	includeForLoops   bool
	// visitor fields
	sliceDeclarations   []*sliceDeclaration
	preallocHints       []analysis.Diagnostic
	returnsInsideOfLoop bool
	arrayTypes          []string
}

func Check(files []*ast.File, simple, includeRangeLoops, includeForLoops bool) []analysis.Diagnostic {
	var hints []analysis.Diagnostic
	for _, f := range files {
		retVis := &returnsVisitor{
			simple:            simple,
			includeRangeLoops: includeRangeLoops,
			includeForLoops:   includeForLoops,
		}
		ast.Walk(retVis, f)
		hints = append(hints, retVis.preallocHints...)
	}

	return hints
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}

	return false
}

func (v *returnsVisitor) Visit(node ast.Node) ast.Visitor {
	v.sliceDeclarations = nil
	v.returnsInsideOfLoop = false
	origLen := len(v.preallocHints)

	switch n := node.(type) {
	case *ast.TypeSpec:
		if _, ok := n.Type.(*ast.ArrayType); ok {
			if n.Name != nil {
				v.arrayTypes = append(v.arrayTypes, n.Name.Name)
			}
		}
	case *ast.BlockStmt:
		if n.List != nil {
			for _, stmt := range n.List {
				switch s := stmt.(type) {
				// Find non pre-allocated slices
				case *ast.DeclStmt:
					genD, ok := s.Decl.(*ast.GenDecl)
					if !ok {
						continue
					}
					if genD.Tok == token.TYPE {
						for _, spec := range genD.Specs {
							tSpec, ok := spec.(*ast.TypeSpec)
							if !ok {
								continue
							}

							if _, ok := tSpec.Type.(*ast.ArrayType); ok {
								if tSpec.Name != nil {
									v.arrayTypes = append(v.arrayTypes, tSpec.Name.Name)
								}
							}
						}
					} else if genD.Tok == token.VAR {
						for _, spec := range genD.Specs {
							vSpec, ok := spec.(*ast.ValueSpec)
							if !ok {
								continue
							}
							if v.isArrayType(vSpec.Type) {
								if vSpec.Names != nil {
									/*atID, ok := arrayType.Elt.(*ast.Ident)
									if !ok {
										continue
									}*/

									// We should handle multiple slices declared on the same line, e.g. var mySlice1, mySlice2 []uint32
									for _, vName := range vSpec.Names {
										v.sliceDeclarations = append(v.sliceDeclarations, &sliceDeclaration{name: vName.Name, pos: genD.Pos()})
									}
								}
							} else if len(vSpec.Names) == len(vSpec.Values) {
								for i, val := range vSpec.Values {
									if v.isCreateEmptyArray(val) {
										v.sliceDeclarations = append(v.sliceDeclarations, &sliceDeclaration{name: vSpec.Names[i].Name, pos: s.Pos()})
									}
								}
							}
						}
					}

				case *ast.AssignStmt:
					if len(s.Lhs) == len(s.Rhs) {
						for index := range s.Lhs {
							ident, ok := s.Lhs[index].(*ast.Ident)
							if !ok {
								continue
							}
							if v.isCreateEmptyArray(s.Rhs[index]) {
								v.sliceDeclarations = append(v.sliceDeclarations, &sliceDeclaration{name: ident.Name, pos: s.Pos()})
							}
						}
					}

				case *ast.RangeStmt:
					if v.includeRangeLoops {
						if len(v.sliceDeclarations) == 0 {
							continue
						}
						// Check the value being ranged over and ensure it's not a channel or an iterator function.
						switch inferExprType(s.X).(type) {
						case *ast.ChanType, *ast.FuncType:
							continue
						}
						if s.Body != nil {
							v.handleLoops(s.Body)
						}
					}

				case *ast.ForStmt:
					if v.includeForLoops {
						if len(v.sliceDeclarations) == 0 {
							continue
						}
						if s.Body != nil {
							v.handleLoops(s.Body)
						}
					}

				default:
				}

				// If simple is true and we had returns inside our loop then discard hints and exit.
				if v.simple && v.returnsInsideOfLoop {
					v.preallocHints = v.preallocHints[:origLen]
					return v
				}
			}
		}
	}
	return v
}

func (v *returnsVisitor) isCreateEmptyArray(expr ast.Expr) bool {
	switch e := expr.(type) {
	case *ast.CompositeLit:
		// []any{}
		return len(e.Elts) == 0 && v.isArrayType(e.Type)
	case *ast.CallExpr:
		switch len(e.Args) {
		case 1:
			// []any(nil)
			arg, ok := e.Args[0].(*ast.Ident)
			if !ok || arg.Name != "nil" {
				return false
			}
			return v.isArrayType(e.Fun)
		case 2:
			// make([]any, 0)
			ident, ok := e.Fun.(*ast.Ident)
			if !ok || ident.Name != "make" {
				return false
			}
			arg, ok := e.Args[1].(*ast.BasicLit)
			if !ok || arg.Value != "0" {
				return false
			}
			return v.isArrayType(e.Args[0])
		}
	}
	return false
}

func (v *returnsVisitor) isArrayType(expr ast.Expr) bool {
	switch e := expr.(type) {
	case *ast.ArrayType:
		return true
	case *ast.Ident:
		return contains(v.arrayTypes, e.Name)
	default:
		return false
	}
}

// handleLoops is a helper function to share the logic required for both *ast.RangeLoops and *ast.ForLoops
func (v *returnsVisitor) handleLoops(blockStmt *ast.BlockStmt) {
	for _, stmt := range blockStmt.List {
		switch bodyStmt := stmt.(type) {
		case *ast.AssignStmt:
			asgnStmt := bodyStmt
			for index, expr := range asgnStmt.Rhs {
				if index >= len(asgnStmt.Lhs) {
					continue
				}

				lhsIdent, ok := asgnStmt.Lhs[index].(*ast.Ident)
				if !ok {
					continue
				}

				callExpr, ok := expr.(*ast.CallExpr)
				if !ok {
					continue
				}

				rhsFuncIdent, ok := callExpr.Fun.(*ast.Ident)
				if !ok {
					continue
				}

				if rhsFuncIdent.Name != "append" {
					continue
				}

				// e.g., `x = append(x)`
				// Pointless, but pre-allocation will not help.
				if len(callExpr.Args) < 2 {
					continue
				}

				rhsIdent, ok := callExpr.Args[0].(*ast.Ident)
				if !ok {
					continue
				}

				// e.g., `x = append(y, a)`
				// This is weird (and maybe a logic error),
				// but we cannot recommend pre-allocation.
				if lhsIdent.Name != rhsIdent.Name {
					continue
				}

				// e.g., `x = append(x, y...)`
				// we should ignore this. Pre-allocating in this case
				// is confusing and is not possible in general.
				if callExpr.Ellipsis.IsValid() {
					continue
				}

				for _, sliceDecl := range v.sliceDeclarations {
					if sliceDecl.name == lhsIdent.Name {
						// This is a potential mark, we just need to make sure there are no returns/continues in the
						// range loop.
						// now we just need to grab whatever we're ranging over
						/*sxIdent, ok := s.X.(*ast.Ident)
						if !ok {
							continue
						}*/

						v.preallocHints = append(v.preallocHints, analysis.Diagnostic{
							Pos:     sliceDecl.pos,
							Message: fmt.Sprintf("Consider preallocating %s", sliceDecl.name),
						})
					}
				}
			}
		case *ast.IfStmt:
			ifStmt := bodyStmt
			if ifStmt.Body != nil {
				for _, ifBodyStmt := range ifStmt.Body.List {
					// TODO: should probably handle embedded ifs here
					switch /*ift :=*/ ifBodyStmt.(type) {
					case *ast.BranchStmt, *ast.ReturnStmt:
						v.returnsInsideOfLoop = true
					default:
					}
				}
			}

		default:
		}
	}
}
