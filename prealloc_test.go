package main

import "testing"

func Test_checkForPreallocations(t *testing.T) {
	const filename = "testdata/sample.go"

	got, err := checkForPreallocations([]string{filename}, true, true, true)
	if err != nil {
		t.Fatal(err)
	}

	want := []Hint{
		Hint{
			LineNumber:        5,
			DeclaredSliceName: "y",
		},
		Hint{
			LineNumber:        6,
			DeclaredSliceName: "z",
		},
		Hint{
			LineNumber:        7,
			DeclaredSliceName: "t",
		},
	}

	if len(got) != len(want) {
		t.Fatalf("expected %d hints, but got %d: %+v", len(want), len(got), got)
	}

	for i := range got {
		act, exp := got[i], want[i]

		if act.Filename != filename {
			t.Errorf("wrong hints[%d].Filename: %q (expected: %q)", i, act.Filename, filename)
		}

		if act.LineNumber != exp.LineNumber {
			t.Errorf("wrong hints[%d].LineNumber: %d (expected: %d)", i, act.LineNumber, exp.LineNumber)
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
