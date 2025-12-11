package test

func returnBeforeAppend() {
	if true {
		return
	}
	var x []int // want "Consider preallocating x"
	x = append(x, 0)
}

func returnAfterAppend() {
	var x []int // want "Consider preallocating x"
	x = append(x, 0)
	if true {
		return
	}
}

func returnBeforeAndAfterAppend() {
	if true {
		return
	}
	var x []int // want "Consider preallocating x"
	x = append(x, 0)
	if true {
		return
	}
}

func returnBetweenAppends() {
	var x []int
	x = append(x, 0)
	if true {
		return
	}
	x = append(x, 0)
}

func gotoAnywhere() {
	var x []int
retry:
	x = append(x, 0)
	if true {
		goto retry
	}
}

func gotoFuncLit() {
	var x []int // want "Consider preallocating x"
	x = append(x, 0)
	f := func() {
		var x []int
	retry:
		x = append(x, 0)
		if true {
			goto retry
		}
	}
	f()
}

func breakLoop() {
	var x []int
	x = append(x, 0)
	for range "Hello" {
		x = append(x, 0)
		break
	}
}

func breakLoopWithoutAppend() {
	var x []int // want "Consider preallocating x"
	x = append(x, 0)
	for range "Hello" {
		break
	}
}

func breakLoopConditional() {
	var x []int
	for i := range "Hello" {
		if true {
			break
		}
		x = append(x, i)
	}
}

func breakLoopSwitch() {
	var x []int // want "Consider preallocating x"
	for range "Hello" {
		switch 0 {
		case 0:
			break
		}
		x = append(x, 0)
	}
}

func breakLoopTypeSwitch() {
	var x []int // want "Consider preallocating x"
	for range "Hello" {
		switch any(x).(type) {
		case []int:
			break
		}
		x = append(x, 0)
	}
}

func breakLoopSelect() {
	var x []int // want "Consider preallocating x"
	for range "Hello" {
		var c chan int
		select {
		case <-c:
			break
		}
		x = append(x, 0)
	}
}

func continueLoop() {
	var x []int
	x = append(x, 0)
	for range "Hello" {
		x = append(x, 0)
		continue
	}
}

func continueLoopWithoutAppend() {
	var x []int // want "Consider preallocating x"
	x = append(x, 0)
	for range "Hello" {
		continue
	}
}

func continueLoopConditional() {
	var x []int
	for i := range "Hello" {
		if true {
			continue
		}
		x = append(x, i)
	}
}
