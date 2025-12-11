package test

// nested statement blocks should be processed to any depth

func nestedBlocks() {
	{
		var x []int // want "Consider preallocating x"
		for i := range "Hello" {
			x = append(x, i)
		}

		if true {
			var y []int // want "Consider preallocating y"
			for i := range "Hello" {
				y = append(y, i)
			}

			for {
				var z []int // want "Consider preallocating z"
				for i := range "Hello" {
					z = append(z, i)
				}
				break
			}
		}
	}
}

func nestedLoops() {
	var x []int // want "Consider preallocating x"
	for i := range "Hello" {
		for j := range "Hello" {
			x = append(x, i, j)
		}
	}
}
