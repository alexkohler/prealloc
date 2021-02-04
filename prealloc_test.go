package main

import (
	"go/token"
	"testing"

	"github.com/alexkohler/prealloc/pkg"
)

func Test_checkForPreallocations(t *testing.T) {
	const filename = "testdata/sample.go"

	fset := token.NewFileSet()

	got, err := checkForPreallocations([]string{filename}, fset, true, true, true)
	if err != nil {
		t.Fatal(err)
	}

	want := []pkg.Hint{
		pkg.Hint{
			Pos:               63,
			DeclaredSliceName: "y",
		},
		pkg.Hint{
			Pos:               77,
			DeclaredSliceName: "z",
		},
		pkg.Hint{
			Pos:               102,
			DeclaredSliceName: "t",
		},
	}

	if len(got) != len(want) {
		t.Fatalf("expected %d hints, but got %d: %+v", len(want), len(got), got)
	}

	for i := range got {
		act, exp := got[i], want[i]

		file := fset.File(act.Pos)

		if file.Name() != filename {
			t.Errorf("wrong hints[%d].Filename: %q (expected: %q)", i, file.Name(), filename)
		}

		actLineNumber := file.Position(act.Pos).Line
		expLineNumber := file.Position(exp.Pos).Line

		if actLineNumber != expLineNumber {
			t.Errorf("wrong hints[%d].LineNumber: %d (expected: %d)", i, actLineNumber, expLineNumber)
		}

		if act.DeclaredSliceName != exp.DeclaredSliceName {
			t.Errorf("wrong hints[%d].DeclaredSliceName: %q (expected: %q)", i, act.DeclaredSliceName, exp.DeclaredSliceName)
		}
	}
}

func BenchmarkSize10NoPreallocate(b *testing.B) {
	existing := make([]int64, 10, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Don't preallocate our initial slice
		var init []int64
		for _, element := range existing {
			init = append(init, element)
		}
	}
}

func BenchmarkSize10Preallocate(b *testing.B) {
	existing := make([]int64, 10, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Preallocate our initial slice
		init := make([]int64, 0, len(existing))
		for _, element := range existing {
			init = append(init, element)
		}
	}
}

func BenchmarkSize10PreallocateCopy(b *testing.B) {
	existing := make([]int64, 10, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Preallocate our initial slice
		init := make([]int64, len(existing))
		copy(init, existing)
	}
}

func BenchmarkSize200NoPreallocate(b *testing.B) {
	existing := make([]int64, 200, 200)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Don't preallocate our initial slice
		var init []int64
		for _, element := range existing {
			init = append(init, element)
		}
	}
}

func BenchmarkSize200Preallocate(b *testing.B) {
	existing := make([]int64, 200, 200)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Preallocate our initial slice
		init := make([]int64, 0, len(existing))
		for _, element := range existing {
			init = append(init, element)
		}
	}
}

func BenchmarkSize200PreallocateCopy(b *testing.B) {
	existing := make([]int64, 200, 200)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Preallocate our initial slice
		init := make([]int64, len(existing))
		copy(init, existing)
	}
}
