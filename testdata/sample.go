package test

func forSimple() {
	var x []int // want "Consider preallocating x"
	for i := 0; i < 5; i++ {
		x = append(x, i)
	}
}

func sliceAssignEmptyLit() {
	x := []int{} // want "Consider preallocating x"
	for i := range "Hello" {
		x = append(x, i)
	}
}

func sliceAssignEmptyMake() {
	x := make([]int, 0) // want "Consider preallocating x"
	for i := range "Hello" {
		x = append(x, i)
	}
}

func sliceAssignNilConvert() {
	x := []int(nil) // want "Consider preallocating x"
	for i := range "Hello" {
		x = append(x, i)
	}
}

func sliceVarAssignEmptyLit() {
	var x = []int{} // want "Consider preallocating x"
	for i := range "Hello" {
		x = append(x, i)
	}
}

func sliceVarAssignEmptyMake() {
	var x = make([]int, 0) // want "Consider preallocating x"
	for i := range "Hello" {
		x = append(x, i)
	}
}

func sliceVarAssignNilConvert() {
	var x = []int(nil) // want "Consider preallocating x"
	for i := range "Hello" {
		x = append(x, i)
	}
}

func sliceAlreadyInitialized() {
	x := []int{1, 2, 3}
	for i := range "Hello" {
		x = append(x, i)
	}
}

func sliceAlreadyAllocated() {
	x := make([]int, 5)
	for i := range "Hello" {
		x = append(x, i)
	}
}

func breakInsideLoop() {
	var x []int
	for i := range "Hello" {
		if true {
			break
		}
		x = append(x, i)
	}
}
