package test

import "iter"

// cannot pre-allocate when ranging over iterators

func rangeSeq() {
	var seq iter.Seq[int]
	var x []int
	for i := range seq {
		x = append(x, i)
	}
}

func rangeSeqArg(seq iter.Seq[int]) {
	var x []int
	for i := range seq {
		x = append(x, i)
	}
}

func rangeSeq2() {
	var seq iter.Seq2[int, int]
	var x []int
	for i := range seq {
		x = append(x, i)
	}
}

func rangeSeq2Arg(seq iter.Seq2[int, int]) {
	var x []int
	for i := range seq {
		x = append(x, i)
	}
}

func rangeFunc() {
	var seq func(func(int) bool)
	var x []int
	for i := range seq {
		x = append(x, i)
	}
}

func rangeFuncArg(seq func(func(int) bool)) {
	var x []int
	for i := range seq {
		x = append(x, i)
	}
}

func rangeFunc2() {
	var seq func(func(int, int) bool)
	var x []int
	for i := range seq {
		x = append(x, i)
	}
}

func rangeFunc2Arg(seq func(func(int, int) bool)) {
	var x []int
	for i := range seq {
		x = append(x, i)
	}
}
