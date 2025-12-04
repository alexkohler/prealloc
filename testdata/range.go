package test

func rangeInt() {
	var x []int // want "Consider preallocating x"
	for i := range 5 {
		x = append(x, i)
	}
}

func rangeIntArg(n int) {
	var x []int // want "Consider preallocating x"
	for i := range n {
		x = append(x, i)
	}
}

func rangeString() {
	var x []int // want "Consider preallocating x"
	for i := range "Hello" {
		x = append(x, i)
	}
}

func rangeStringArg(s string) {
	var x []int // want "Consider preallocating x"
	for i := range s {
		x = append(x, i)
	}
}

func rangeTwice() {
	var x []int // want "Consider preallocating x"
	for i := range 5 {
		x = append(x, i)
	}
	for i := range "Hello" {
		x = append(x, i)
	}
}
