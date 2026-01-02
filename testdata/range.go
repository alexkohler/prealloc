package test

import "sort"

func rangeZero() {
	var x []int
	for i := range 0 {
		x = append(x, i)
	}
}

func rangeEmptyString() {
	var x []int
	for i := range "" {
		x = append(x, i)
	}
}

func rangeInt() {
	var x []int // want "Consider preallocating x with capacity 5$"
	for i := range 5 {
		x = append(x, i)
	}
}

func rangeIntVar() {
	n := 5
	var x []int // want "Consider preallocating x with capacity n$"
	for i := range n {
		x = append(x, i)
	}
}

func rangeIntArg(n int) {
	var x []int // want "Consider preallocating x with capacity n$"
	for i := range n {
		x = append(x, i)
	}
}

func rangeString() {
	var x []int // want "Consider preallocating x with capacity 5$"
	for i := range "Hello" {
		x = append(x, i)
	}
}

func rangeIntFunc() {
	fn := func() int { return 5 }
	var x []int // want "Consider preallocating x$"
	for i := range fn() {
		x = append(x, i)
	}
}

func rangeStringVar() {
	s := "Hello"
	var x []int // want "Consider preallocating x with capacity len\\(s\\)$"
	for i := range s {
		x = append(x, i)
	}
}

func rangeStringArg(s string) {
	var x []int // want "Consider preallocating x with capacity len\\(s\\)$"
	for i := range s {
		x = append(x, i)
	}
}

func rangeStringFunc() {
	fn := func() string { return "Hello" }
	var x []int // want "Consider preallocating x$"
	for i := range fn() {
		x = append(x, i)
	}
}

func rangeSliceVar() {
	var a []int
	var x []int // want "Consider preallocating x with capacity len\\(a\\)$"
	for i := range a {
		x = append(x, i)
	}
}

func rangeSliceLit() {
	var x []int // want "Consider preallocating x with capacity 3$"
	for i := range []int{1, 2, 3} {
		x = append(x, i)
	}
}

func rangeSliceTypeConvert() {
	var x []int // want "Consider preallocating x with capacity 5$"
	for i := range []byte("Hello") {
		x = append(x, i)
	}
}

func rangeArrayVar() {
	var a [5]int
	var x []int // want "Consider preallocating x with capacity len\\(a\\)$"
	for i := range a {
		x = append(x, i)
	}
}

func rangeSliceFunc() {
	fn := func() []int { return []int{1, 2, 3} }
	var x []int // want "Consider preallocating x$"
	for i := range fn() {
		x = append(x, i)
	}
}

func rangeArrayLit() {
	var x []int // want "Consider preallocating x with capacity 5$"
	for i := range [5]int{1, 2, 3} {
		x = append(x, i)
	}
}

func rangeArrayFunc() {
	fn := func() [5]int { return [5]int{1, 2, 3} }
	var x []int // want "Consider preallocating x$"
	for i := range fn() {
		x = append(x, i)
	}
}

func rangeArrayPointerVar() {
	var a *[5]int
	var x []int // want "Consider preallocating x with capacity len\\(a\\)$"
	for i := range a {
		x = append(x, i)
	}
}

func rangeArrayPointerLit() {
	var x []int // want "Consider preallocating x with capacity 5$"
	for i := range &[5]int{1, 2, 3} {
		x = append(x, i)
	}
}

func rangeArrayPointerFunc() {
	fn := func() *[5]int { return &[5]int{1, 2, 3} }
	var x []int // want "Consider preallocating x$"
	for i := range fn() {
		x = append(x, i)
	}
}

func rangeMapVar() {
	var m map[int]int
	var x []int // want "Consider preallocating x with capacity len\\(m\\)$"
	for i := range m {
		x = append(x, i)
	}
}

func rangeMapLit() {
	var x []int // want "Consider preallocating x with capacity 2$"
	for i := range map[int]int{1: 2, 3: 4} {
		x = append(x, i)
	}
}

func rangeMapFunc() {
	fn := func() map[int]int { return map[int]int{1: 2, 3: 4} }
	var x []int // want "Consider preallocating x$"
	for i := range fn() {
		x = append(x, i)
	}
}

func rangeIntTypeConvert() {
	var x []uint // want "Consider preallocating x with capacity 5$"
	for i := range uint(5) {
		x = append(x, i)
	}
}

func rangeMultiple() {
	var x []int // want "Consider preallocating x with capacity 5 \\+ n \\+ len\\(s\\) \\+ \\(n - m \\+ 1\\)$"
	for i := range 5 {
		x = append(x, i)
	}
	n := 5
	for i := range n {
		x = append(x, i)
	}
	s := "Hello"
	for i := range s {
		x = append(x, i)
	}
	m := 0
	for i := m; i <= n; i++ {
		x = append(x, i)
	}
}

func rangeMultipleWithPartialUnresolvedCapacity() {
	var x []int // want "Consider preallocating x$"
	for i := range 5 {
		x = append(x, i)
	}
	var s sort.IntSlice
	for i := range s {
		x = append(x, i)
	}
}

func rangeSliceReslice() {
	var a []int
	var x []int // want "Consider preallocating x with capacity len\\(a\\)$"
	for i := range a[:] {
		x = append(x, i)
	}
}

func rangeSlicePrefix() {
	var a []int
	var x []int // want "Consider preallocating x with capacity 4$"
	for i := range a[:4] {
		x = append(x, i)
	}
}

func rangeSliceSubslice() {
	var a []int
	var x []int // want "Consider preallocating x with capacity 2$"
	for i := range a[2:4] {
		x = append(x, i)
	}
}

func rangeSliceSuffix() {
	var a []int
	var x []int // want "Consider preallocating x with capacity len\\(a\\) - 2$"
	for i := range a[2:] {
		x = append(x, i)
	}
}

func rangeAppendLit() {
	var x []int // want "Consider preallocating x with capacity 4$"
	for i := range append([]int{1, 2}, 3, 4) {
		x = append(x, i)
	}
}

func rangeAppendLitEllipsisLit() {
	var x []int // want "Consider preallocating x with capacity 4$"
	for i := range append([]int{1, 2}, []int{3, 4}...) {
		x = append(x, i)
	}
}

func rangeAppendVar() {
	var y []int
	var x []int // want "Consider preallocating x with capacity len\\(y\\) \\+ 2$"
	for i := range append(y, 3, 4) {
		x = append(x, i)
	}
}

func rangeAppendLitEllipsisVar() {
	var y []int
	var x []int // want "Consider preallocating x with capacity 2 \\+ len\\(y\\)$"
	for i := range append([]int{1, 2}, y...) {
		x = append(x, i)
	}
}
