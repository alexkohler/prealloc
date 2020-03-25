package main

func main() {
	x := make([]rune, len("Hello"))
	var y []rune
	var z, w, v, u []int
	var t [][]int

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
	}

	_ = v
}

func foo(n int) []int {
	return make([]int, n)
}
