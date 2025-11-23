package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestCheckForPreallocations(t *testing.T) {
	t.Parallel()

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %v", err)
	}

	analysistest.Run(t, filepath.Join(wd, "testdata"), NewAnalyzer(), ".")
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
