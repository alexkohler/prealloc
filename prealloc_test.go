package main

import (
	"fmt"
	"go/token"
	"testing"

	"github.com/alexkohler/prealloc/pkg"
)

func TestCheckForPreallocations(t *testing.T) {
	t.Parallel()

	const filename = "testdata/sample.go"

	fset := token.NewFileSet()

	got, err := checkForPreallocations([]string{filename}, fset, true, true, true)
	if err != nil {
		t.Fatal(err)
	}

	want := []pkg.Hint{
		{
			Pos:               78,
			DeclaredSliceName: "y",
		},
		{
			Pos:               92,
			DeclaredSliceName: "z",
		},
		{
			Pos:               117,
			DeclaredSliceName: "t",
		},
		{
			Pos:               183,
			DeclaredSliceName: "a",
		},
		{
			Pos:               218,
			DeclaredSliceName: "c",
		},
		{
			Pos:               247,
			DeclaredSliceName: "e",
		},
		{
			Pos:               295,
			DeclaredSliceName: "g",
		},
		{
			Pos:               337,
			DeclaredSliceName: "j",
		},
		{
			Pos:               382,
			DeclaredSliceName: "l",
		},
		{
			Pos:               2551,
			DeclaredSliceName: "m",
		},
		{
			Pos:               2671,
			DeclaredSliceName: "n",
		},
		{
			Pos:               2793,
			DeclaredSliceName: "o",
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
	existing := make([]int64, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Don't preallocate our initial slice
		var init []int64
		for _, element := range existing { //nolint:staticcheck
			init = append(init, element) //nolint:staticcheck
		}
	}
}

func BenchmarkSize10Preallocate(b *testing.B) {
	existing := make([]int64, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Preallocate our initial slice
		init := make([]int64, 0, len(existing))
		for _, element := range existing { //nolint:staticcheck
			init = append(init, element) //nolint:staticcheck
		}
	}
}

func BenchmarkSize10PreallocateCopy(b *testing.B) {
	existing := make([]int64, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Preallocate our initial slice
		init := make([]int64, len(existing))
		copy(init, existing)
	}
}

func BenchmarkSize200NoPreallocate(b *testing.B) {
	existing := make([]int64, 200)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Don't preallocate our initial slice
		var init []int64
		for _, element := range existing { //nolint:staticcheck
			init = append(init, element) //nolint:staticcheck
		}
	}
}

func BenchmarkSize200Preallocate(b *testing.B) {
	existing := make([]int64, 200)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Preallocate our initial slice
		init := make([]int64, 0, len(existing))
		for _, element := range existing { //nolint:staticcheck
			init = append(init, element) //nolint:staticcheck
		}
	}
}

func BenchmarkSize200PreallocateCopy(b *testing.B) {
	existing := make([]int64, 200)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Preallocate our initial slice
		init := make([]int64, len(existing))
		copy(init, existing)
	}
}

func BenchmarkMap(b *testing.B) {
	benchmarks := []struct {
		size        int
		preallocate bool
	}{
		{10, false},
		{10, true},
		{200, false},
		{200, true},
	}
	var m map[int]int
	for _, bm := range benchmarks {
		no := ""
		if !bm.preallocate {
			no = "No"
		}
		b.Run(fmt.Sprintf("Size%d%sPreallocate", bm.size, no), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if bm.preallocate {
					m = make(map[int]int, bm.size)
				} else {
					m = make(map[int]int)
				}
				for j := 0; j < bm.size; j++ {
					m[j] = j
				}
			}
		})
	}
}
