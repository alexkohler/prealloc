package main

import "testing"

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
