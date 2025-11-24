package test

// cannot pre-allocate when ranging over channels

func rangeChan() {
	var ch chan int
	var x []int
	for i := range ch {
		x = append(x, i)
	}
}

func rangeChanMake() {
	ch := make(chan int)
	var x []int
	for i := range ch {
		x = append(x, i)
	}
}

func rangeChanArg(ch chan int) {
	var x []int
	for i := range ch {
		x = append(x, i)
	}
}
