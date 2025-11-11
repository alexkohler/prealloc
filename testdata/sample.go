package main

func main() {
	x := make([]rune, len("Hello"))
	var y []rune
	var z, w, v, u, s []int
	var t [][]int
	var intChan chan int
	var a = make([]int, 0)
	var b = make([]int, 1)
	c := make([]int, 0)
	d := make([]int, 1)

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

		// b is not a candidate for pre-allocation since it was initialized non-empty
		b = append(b, i)

		// c is a candidate for pre-allocation
		c = append(c, i)

		// d is not a candidate for pre-allocation since it was initialized non-empty
		d = append(d, i)
	}

	for i := range intChan {
		// s is not a candidate for pre-allocation since the range target is a channel
		s = append(s, i)
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
