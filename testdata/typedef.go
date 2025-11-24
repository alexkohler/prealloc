package test

import . "sort"

type ints []int

func assignTypeDefEmptyLit() {
	x := ints{} // want "Consider preallocating x"
	for i := range "Hello" {
		x = append(x, i)
	}
}

func assignTypeDefEmptyMake() {
	x := make(ints, 0) // want "Consider preallocating x"
	for i := range "Hello" {
		x = append(x, i)
	}
}

func assignTypeDefNilConvert() {
	x := ints(nil) // want "Consider preallocating x"
	for i := range 5 {
		x = append(x, i)
	}
}

func varAssignTypeDefEmptyLit() {
	var x = ints{} // want "Consider preallocating x"
	for i := range 5 {
		x = append(x, i)
	}
}

func varAssignTypeDefEmptyMake() {
	var x = make(ints, 0) // want "Consider preallocating x"
	for i := range "Hello" {
		x = append(x, i)
	}
}

func varAssignTypeDefNilConvert() {
	var x = ints(nil) // want "Consider preallocating x"
	for i := range 5 {
		x = append(x, i)
	}
}

func inlineTypeDefEmptyLit() {
	type ints []int
	var x ints // want "Consider preallocating x"
	for i := range "Hello" {
		x = append(x, i)
	}
}

func externalTypeDefEmptyLit() {
	var x IntSlice
	for i := range "Hello" {
		x = append(x, i)
	}
}
