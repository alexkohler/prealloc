package pkg

import (
	"fmt"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

type sliceDeclaration struct {
	name      string
	declPos   token.Pos
	level     int       // Nesting level of this slice. Will be disqualified if appended at a deeper level.
	appendPos token.Pos // Position of most recent append. Used to determine if appended in an unsupported loop.
	exclude   bool      // Whether this slice has been disqualified due to an unsupported pattern.
	hasReturn bool      // Whether a return statement has been found after the first append. Any subsequent appends will disqualify this slice in simple mode.
}

type returnsVisitor struct {
	// flags
	simple            bool
	includeRangeLoops bool
	includeForLoops   bool
	// visitor fields
	sliceDeclarations []*sliceDeclaration
	preallocHints     []analysis.Diagnostic
	level             int  // Current nesting level. Loops do not increment the level.
	hasReturn         bool // Whether a return statement has been found. Slices appended before and after a return are disqualified in simple mode.
	hasGoto           bool // Whether a goto statement has been found. Goto disqualifies pending and subsequent slices in simple mode.
	hasBranch         bool // Whether a branch statement has been found. Loops with branch statements are unsupported in simple mode.
}

func Check(files []*ast.File, simple, includeRangeLoops, includeForLoops bool) []analysis.Diagnostic {
	retVis := &returnsVisitor{
		simple:            simple,
		includeRangeLoops: includeRangeLoops,
		includeForLoops:   includeForLoops,
	}
	for _, f := range files {
		ast.Walk(retVis, f)
	}
	return retVis.preallocHints
}

func (v *returnsVisitor) Visit(node ast.Node) ast.Visitor {
	switch s := node.(type) {
	case *ast.FuncDecl:
		ast.Walk(v, s.Body)
		v.level = 0
		v.hasReturn = false
		v.hasGoto = false
		return nil

	case *ast.FuncLit:
		wasReturn := v.hasReturn
		wasGoto := v.hasGoto
		v.hasReturn = false
		ast.Walk(v, s.Body)
		v.hasReturn = wasReturn
		v.hasGoto = wasGoto
		return nil

	case *ast.BlockStmt:
		declIdx := len(v.sliceDeclarations)
		v.level++
		for _, stmt := range s.List {
			ast.Walk(v, stmt)
		}
		v.level--
		for i := declIdx; i < len(v.sliceDeclarations); i++ {
			sliceDecl := v.sliceDeclarations[i]
			if sliceDecl.appendPos.IsValid() && !sliceDecl.exclude && !v.hasGoto {
				v.preallocHints = append(v.preallocHints, analysis.Diagnostic{
					Pos:     sliceDecl.declPos,
					Message: fmt.Sprintf("Consider preallocating %s", sliceDecl.name),
				})
			}
		}
		v.sliceDeclarations = v.sliceDeclarations[:declIdx]
		return nil

	case *ast.DeclStmt:
		genD, ok := s.Decl.(*ast.GenDecl)
		if !ok || genD.Tok != token.VAR {
			return nil
		}
		for _, spec := range genD.Specs {
			vSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			if len(vSpec.Values) == 0 {
				if _, ok := inferExprType(vSpec.Type).(*ast.ArrayType); ok {
					for _, vName := range vSpec.Names {
						v.sliceDeclarations = append(v.sliceDeclarations, &sliceDeclaration{name: vName.Name, declPos: s.Pos(), level: v.level})
					}
				}
			} else {
				for i, vName := range vSpec.Names {
					if i < len(vSpec.Values) && isCreateEmptyArray(vSpec.Values[i]) {
						v.sliceDeclarations = append(v.sliceDeclarations, &sliceDeclaration{name: vName.Name, declPos: s.Pos(), level: v.level})
					}
				}
			}
		}

	case *ast.AssignStmt:
		if len(s.Lhs) != len(s.Rhs) {
			return nil
		}
		for i, lhs := range s.Lhs {
			ident, ok := lhs.(*ast.Ident)
			if !ok {
				continue
			}
			if isCreateEmptyArray(s.Rhs[i]) {
				v.sliceDeclarations = append(v.sliceDeclarations, &sliceDeclaration{name: ident.Name, declPos: s.Pos(), level: v.level})
			} else {
				switch expr := s.Rhs[i].(type) {
				case *ast.Ident:
					if s.Tok != token.ASSIGN || expr.Name != "nil" {
						continue
					}
					for _, sliceDecl := range v.sliceDeclarations {
						if sliceDecl.name == ident.Name {
							v.sliceDeclarations = append(v.sliceDeclarations, &sliceDeclaration{name: ident.Name, declPos: s.Pos(), level: v.level})
							break
						}
					}
				case *ast.CallExpr:
					if len(expr.Args) < 2 {
						continue
					}
					if funIdent, ok := expr.Fun.(*ast.Ident); !ok || funIdent.Name != "append" {
						continue
					}
					rhsIdent, ok := expr.Args[0].(*ast.Ident)
					if !ok {
						continue
					}
					for i := len(v.sliceDeclarations) - 1; i >= 0; i-- {
						sliceDecl := v.sliceDeclarations[i]
						if sliceDecl.name == ident.Name {
							if expr.Ellipsis.IsValid() || ident.Name != rhsIdent.Name || sliceDecl.hasReturn || sliceDecl.level != v.level {
								sliceDecl.exclude = true
							} else {
								sliceDecl.appendPos = s.Pos()
							}
							break
						}
					}
				}
			}
		}

	case *ast.RangeStmt:
		if len(v.sliceDeclarations) == 0 {
			return v
		}
		hadBranch := v.hasBranch
		v.hasBranch = false
		v.level--
		ast.Walk(v, s.Body)
		v.level++
		exclude := !v.includeRangeLoops || v.hasReturn || v.hasGoto || v.hasBranch
		if !exclude {
			switch inferExprType(s.X).(type) {
			case *ast.ChanType, *ast.FuncType:
				exclude = true
			}
		}
		if exclude {
			// exclude all slices that were appended within this loop
			for _, sliceDecl := range v.sliceDeclarations {
				if sliceDecl.appendPos > s.Pos() {
					sliceDecl.exclude = true
				}
			}
		}
		v.hasBranch = hadBranch
		return nil

	case *ast.ForStmt:
		if len(v.sliceDeclarations) == 0 {
			return v
		}
		hadBranch := v.hasBranch
		v.hasBranch = false
		v.level--
		ast.Walk(v, s.Body)
		v.level++
		if !v.includeForLoops || v.hasReturn || v.hasGoto || v.hasBranch || s.Init == nil || s.Cond == nil || s.Post == nil {
			// exclude all slices that were appended within this loop
			for _, sliceDecl := range v.sliceDeclarations {
				if sliceDecl.appendPos > s.Pos() {
					sliceDecl.exclude = true
				}
			}
		}
		v.hasBranch = hadBranch
		return nil

	case *ast.SwitchStmt:
		return v.walkSwitchSelect(s.Body)

	case *ast.TypeSwitchStmt:
		return v.walkSwitchSelect(s.Body)

	case *ast.SelectStmt:
		return v.walkSwitchSelect(s.Body)

	case *ast.ReturnStmt:
		if !v.simple {
			return nil
		}
		v.hasReturn = true
		// flag all slices that have been appended at least once
		for _, sliceDecl := range v.sliceDeclarations {
			if sliceDecl.appendPos.IsValid() {
				sliceDecl.hasReturn = true
			}
		}

	case *ast.BranchStmt:
		if !v.simple {
			return nil
		}
		if s.Label != nil {
			v.hasGoto = true
		} else {
			v.hasBranch = true
		}
	}

	return v
}

func (v *returnsVisitor) walkSwitchSelect(body *ast.BlockStmt) ast.Visitor {
	hadBranch := v.hasBranch
	v.hasBranch = false
	ast.Walk(v, body)
	v.hasBranch = hadBranch
	return nil
}

func isCreateEmptyArray(expr ast.Expr) bool {
	switch e := expr.(type) {
	case *ast.CompositeLit:
		// []any{}
		_, ok := inferExprType(e.Type).(*ast.ArrayType)
		return ok && len(e.Elts) == 0
	case *ast.CallExpr:
		switch len(e.Args) {
		case 1:
			// []any(nil)
			arg, ok := e.Args[0].(*ast.Ident)
			if !ok || arg.Name != "nil" {
				return false
			}
			_, ok = inferExprType(e.Fun).(*ast.ArrayType)
			return ok
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
			_, ok = inferExprType(e.Args[0]).(*ast.ArrayType)
			return ok
		}
	}
	return false
}
