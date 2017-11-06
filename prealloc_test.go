package main

import "testing"

func BenchmarkIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		existing := make([]int64, 1000, 1000)
		init := make([]int64, 1000) // len 1000, cap 1000
		for index, element := range existing {
			init[index] = element
		}
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		existing := make([]int64, 1000, 1000)
		var init []int64
		for _, element := range existing {
			init = append(init, element)
		}
	}
}
