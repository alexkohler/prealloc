package pkg

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestInferValueTypeMissingValuesDoesNotPanic(t *testing.T) {
	t.Parallel()

	spec := &ast.ValueSpec{
		Names: []*ast.Ident{{Name: "n"}},
	}

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("inferValueType panicked: %v", r)
		}
	}()

	if got := inferValueType(spec, "n"); got != nil {
		t.Fatalf("inferValueType() = %T, want nil", got)
	}
}

func TestInferAssignTypeMissingRHSDoesNotPanic(t *testing.T) {
	t.Parallel()

	assign := &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent("x")},
		Rhs: nil,
	}

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("inferAssignType panicked: %v", r)
		}
	}()

	if got := inferAssignType(assign, "x"); got != nil {
		t.Fatalf("inferAssignType() = %T, want nil", got)
	}
}

func TestCheckMalformedVarSpecDoesNotPanic(t *testing.T) {
	t.Parallel()

	src := `package test

var n, m =

func f() {
	var out []int
	for i := range n {
		out = append(out, i)
	}
}
`

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "malformed.go", src, parser.AllErrors)
	if file == nil {
		t.Fatalf("ParseFile returned nil file: %v", err)
	}
	if err == nil {
		t.Fatalf("expected parser error for malformed source")
	}

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Check panicked: %v", r)
		}
	}()

	_ = Check([]*ast.File{file}, true, true, false)
}
