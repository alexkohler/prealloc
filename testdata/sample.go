package main

import "iter"

func main() {
	x := make([]rune, len("Hello"))
	var y []rune
	var z, w, v, u, s []int
	var t [][]int
	var intChan chan int
	intChan2 := make(chan int)

	var a = []int{}
	var b = []int{0}
	c := []int{}
	d := []int{0}
	var e = make([]int, 0)
	var f = make([]int, 1)
	g := make([]int, 0)
	h := make([]int, 1)
	var j = []int(nil)
	var k = []int([]int{0})
	l := []int(nil)
	p := []int([]int{0})
	var q, r, m []int

	for i, r := range "Hello" {
		// x is already pre-allocated
		// y is a candidate for pre-allocation
		x[i], y = r, append(y, r)

		// w is not a candidate for pre-allocation due to `...`
		w = append(w, foo(i)...)

		// v is not a candidate for pre-allocation since this appends to u
		v = append(u, i)

		// u is not a candidate for pre-allocation since nothing was actually appended
		u = append(u)

		// z is a candidate for pre-allocation
		z = append(z, i)

		// t is a candidate for pre-allocation
		t = append(t, foo(i))

		// a is a candidate for pre-allocation
		a = append(a, i)
		// b is not a candidate for pre-allocation since it was initialized with values
		b = append(b, i)
		// c is a candidate for pre-allocation
		c = append(c, i)
		// d is not a candidate for pre-allocation since it was initialized with values
		d = append(d, i)
		// e is a candidate for pre-allocation
		e = append(e, i)
		// f is not a candidate for pre-allocation since it was initialized non-empty
		f = append(f, i)
		// g is a candidate for pre-allocation
		g = append(g, i)
		// h is not a candidate for pre-allocation since it was initialized non-empty
		h = append(h, i)
		// j is a candidate for pre-allocation
		j = append(j, i)
		// k is not a candidate for pre-allocation since it was converted from a non-empty slice
		k = append(k, i)
		// l is a candidate for pre-allocation
		l = append(l, i)
		// p is not a candidate for pre-allocation since it was converted from a non-empty slice
		p = append(p, i)
	}

	for i := range intChan {
		// s is not a candidate for pre-allocation since the range target is a channel
		s = append(s, i)
	}

	for i := range intChan2 {
		// q is not a candidate for pre-allocation since the range target is a channel
		q = append(q, i)
	}

	var intSeq iter.Seq[int]
	for i := range intSeq {
		// r is not a candidate for pre-allocation since the range target is an iterator
		r = append(r, i)
	}

	var intSeq2 iter.Seq2[int, int]
	for i := range intSeq2 {
		// m is not a candidate for pre-allocation since the range target is an iterator
		m = append(m, i)
	}

	_ = v

	{
		var m []int
		for i := range "Hello" {
			// m is a candidate for preallocation
			m = append(m, i)
		}

		if true {
			var n []int
			for i := range "Hello" {
				// n is a candidate for preallocation
				n = append(n, i)
			}

			for {
				var o []int
				for i := range "Hello" {
					// o is a candidate for preallocation
					o = append(o, i)
				}
				break
			}
		}
	}
}

func foo(n int) []int {
	return make([]int, n)
}
