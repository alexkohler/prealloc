package test

func appendNothing() {
	var x []int
	for range "Hello" {
		x = append(x)
	}
}

func appendToAnother() {
	var x []int
	var y []int
	x = append(x, 0)
	for i := range "Hello" {
		x = append(y, i)
	}
	_ = x
}

func appendEllipsisVar() {
	var nums []int
	var x []int // want "Consider preallocating x with capacity len\\(nums\\) \\* 5$"
	for range "Hello" {
		x = append(x, nums...)
	}
}

func appendEllipsisLit() {
	var x []int // want "Consider preallocating x with capacity 15$"
	for range "Hello" {
		x = append(x, []int{1, 2, 3}...)
	}
}

func appendEllipsisString() {
	var x []byte // want "Consider preallocating x with capacity 25$"
	for range 5 {
		x = append(x, "hello"...)
	}
}

func appendEllipsisReslice() {
	var y []int
	var x []int // want "Consider preallocating x with capacity len\\(y\\) \\* 5$"
	for range 5 {
		x = append(x, y[:]...)
	}
}

func appendEllipsisPrefix() {
	var y []int
	var x []int // want "Consider preallocating x with capacity 20$"
	for range 5 {
		x = append(x, y[:4]...)
	}
}

func appendEllipsisSubslice() {
	var y []int
	var x []int // want "Consider preallocating x with capacity 10$"
	for range 5 {
		x = append(x, y[2:4]...)
	}
}

func appendEllipsisSuffix() {
	var y []int
	var x []int // want "Consider preallocating x with capacity \\(len\\(y\\) - 2\\) \\* 5$"
	for range 5 {
		x = append(x, y[2:]...)
	}
}

func appendEllipsisFunc() {
	fn := func() []int { return []int{1, 2, 3} }
	var x []int // want "Consider preallocating x$"
	for range 5 {
		x = append(x, fn()...)
	}
}

func appendEllipsisTypeConvert() {
	var x []byte // want "Consider preallocating x with capacity 25$"
	for range 5 {
		x = append(x, []byte("Hello")...)
	}
}

func appendMultipleCalls() {
	var x []int // want "Consider preallocating x with capacity 10$"
	for i := range 5 {
		x = append(x, i)
		x = append(x, i)
	}
}

func appendMultipleArgs() {
	var x []int // want "Consider preallocating x with capacity 10$"
	for i := range 5 {
		x = append(x, i, i)
	}
}

func appendMultipleRangeIntVar() {
	n := 5
	var x []int // want "Consider preallocating x with capacity 2 \\* n$"
	for i := range n {
		x = append(x, i, i)
	}
}

func appendMultipleRangeStringVar() {
	s := "Hello"
	var x []int // want "Consider preallocating x with capacity 2 \\* len\\(s\\)$"
	for i := range s {
		x = append(x, i, i)
	}
}

func appendAppend() {
	s := "Hello"
	var x []int // want "Consider preallocating x with capacity 2 \\* len\\(s\\)$"
	for i := range s {
		x = append(x, append([]int{i}, i)...)
	}
}
