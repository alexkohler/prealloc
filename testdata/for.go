package test

func forSimple() {
	var x []int // want "Consider preallocating x"
	for i := 0; i < 5; i++ {
		x = append(x, i)
	}
}

func forInfinite() {
	var x []int
	for {
		x = append(x, 0)
	}
}

func forWhile() {
	var x []int
	for true {
		x = append(x, 0)
	}
}
