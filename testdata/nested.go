package test

// nested statement blocks should be processed to any depth

func nestedBlocks() {
	{
		var x []int // want "Consider preallocating x with capacity 5$"
		for i := range "Hello" {
			x = append(x, i)
		}

		if true {
			var y []int // want "Consider preallocating y with capacity 5$"
			for i := range "Hello" {
				y = append(y, i)
			}

			for {
				var z []int // want "Consider preallocating z with capacity 5$"
				for i := range "Hello" {
					z = append(z, i)
				}
				break
			}
		}
	}
}

func nestedRangeIndependent() {
	var x []int // want "Consider preallocating x with capacity 50$"
	for i := range "Hello" {
		for j := range "Hello" {
			x = append(x, i, j)
		}
	}
}

func nestedRangeDependentKey() {
	var x []int
	for i := range "Hello" {
		for j := range i {
			x = append(x, i, j)
		}
	}
}

func nestedRangeDependentValue() {
	var x []int
	for i, v := range "Hello" {
		for j := range v {
			x = append(x, i, int(j))
		}
	}
}

func nestedRangeDependentNewVar() {
	var x []int
	for i := range "Hello" {
		ii := i
		for j := range ii {
			x = append(x, ii, j)
		}
	}
}

func nestedForDependentLower() {
	var x []int
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			x = append(x, i, j)
		}
	}
}

func nestedForDependentUpper() {
	var x []int
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			x = append(x, i, j)
		}
	}
}

func nestedReturn() {
	var x []int
	for i := range "Hello" {
		x = append(x, i)
		{
			if true {
				for {
					return
				}
			}
		}
	}
}

func nestedBreak() {
	var x []int
	for i := range "Hello" {
		{
			if true {
				for {
					x = append(x, i)
					break
				}
			}
		}
	}
}

func nestedContinue() {
	var x []int
	for i := range "Hello" {
		{
			if true {
				for {
					x = append(x, i)
					continue
				}
			}
		}
	}
}
