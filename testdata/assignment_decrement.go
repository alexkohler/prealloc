package test

func forAssignmentDecrement() {
	var x []int // want "Consider preallocating x with capacity 5$"
	for i := 5; i > 0; i = i - 1 {
		x = append(x, i)
	}
}

func forAssignmentDecrementStride() {
	var x []int // want "Consider preallocating x with capacity 3$"
	for i := 6; i > 0; i = i - 2 {
		x = append(x, i)
	}
}

func forAssignmentDecrementInclusive() {
	var x []int // want "Consider preallocating x with capacity 6$"
	for i := 5; i >= 0; i = i - 1 {
		x = append(x, i)
	}
}

func forAssignmentReflectedSubtraction() {
	var x []int // want "Consider preallocating x$"
	for i := 0; i < 5; i = 6 - i {
		x = append(x, i)
	}
}
