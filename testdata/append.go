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
	for i := range "Hello" {
		x = append(y, i)
	}
	_ = x
}

func appendEllipsis() {
	var nums []int
	var x []int
	for range "Hello" {
		x = append(x, nums...)
	}
}
