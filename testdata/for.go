package test

import "bytes"

func forInfinite() {
	var x []int
	for {
		x = append(x, 0)
	}
}

func forWhile() {
	var x []int
	for true {
		x = append(x, 0)
	}
}

func forIncZeroToMaxExclusive() {
	var x []int // want "Consider preallocating x with capacity 5$"
	for i := 0; i < 5; i++ {
		x = append(x, i)
	}
}

func forIncOneToMaxExclusive() {
	var x []int // want "Consider preallocating x with capacity 4$"
	for i := 1; i < 5; i++ {
		x = append(x, i)
	}
}

func forIncZeroToMaxInclusive() {
	var x []int // want "Consider preallocating x with capacity 6$"
	for i := 0; i <= 5; i++ {
		x = append(x, i)
	}
}

func forIncZeroToNotMax() {
	var x []int // want "Consider preallocating x with capacity 5$"
	for i := 0; i != 5; i++ {
		x = append(x, i)
	}
}

func forDecMaxToZeroExclusive() {
	var x []int // want "Consider preallocating x with capacity 5$"
	for i := 5; i > 0; i-- {
		x = append(x, i)
	}
}

func forDecMaxToOneExclusive() {
	var x []int // want "Consider preallocating x with capacity 4$"
	for i := 5; i > 1; i-- {
		x = append(x, i)
	}
}

func forDecMaxToZeroInclusive() {
	var x []int // want "Consider preallocating x with capacity 6$"
	for i := 5; i >= 0; i-- {
		x = append(x, i)
	}
}

func forDecMaxToNotZero() {
	var x []int // want "Consider preallocating x with capacity 5$"
	for i := 5; i != 0; i-- {
		x = append(x, i)
	}
}

func forIncZeroToMaxExcReverse() {
	var x []int // want "Consider preallocating x with capacity 5$"
	for i := 0; 5 > i; i++ {
		x = append(x, i)
	}
}

func forIncZeroToMaxIncReverse() {
	var x []int // want "Consider preallocating x with capacity 6$"
	for i := 0; 5 >= i; i++ {
		x = append(x, i)
	}
}

func forDecMaxToZeroExcReverse() {
	var x []int // want "Consider preallocating x with capacity 5$"
	for i := 5; 0 < i; i-- {
		x = append(x, i)
	}
}

func forDecMaxToZeroIncReverse() {
	var x []int // want "Consider preallocating x with capacity 6$"
	for i := 5; 0 <= i; i-- {
		x = append(x, i)
	}
}

func forIncZeroToVarExclusive() {
	n := 5
	var x []int // want "Consider preallocating x with capacity n$"
	for i := 0; i < n; i++ {
		x = append(x, i)
	}
}

func forIncOneToVarExclusive() {
	n := 5
	var x []int // want "Consider preallocating x with capacity n - 1$"
	for i := 1; i < n; i++ {
		x = append(x, i)
	}
}

func forIncVarNegativeOneToVarExclusive() {
	n := 5
	var x []int // want "Consider preallocating x with capacity n \\+ 1$"
	for i := -1; i < n; i++ {
		x = append(x, i)
	}
}

func forIncVarToMaxExclusive() {
	m := 0
	var x []int // want "Consider preallocating x with capacity 5 - m$"
	for i := m; i < 5; i++ {
		x = append(x, i)
	}
}

func forIncVarToMaxInclusive() {
	m := 0
	var x []int // want "Consider preallocating x with capacity 6 - m$"
	for i := m; i <= 5; i++ {
		x = append(x, i)
	}
}

func forIncVarToZeroExclusive() {
	m := -5
	var x []int // want "Consider preallocating x with capacity -m$"
	for i := m; i < 0; i++ {
		x = append(x, i)
	}
}

func forIncVarToVarExclusive() {
	m := 0
	n := 5
	var x []int // want "Consider preallocating x with capacity n - m$"
	for i := m; i < n; i++ {
		x = append(x, i)
	}
}

func forIncVarToVarInclusive() {
	m := 0
	n := 5
	var x []int // want "Consider preallocating x with capacity n \\+ 1 - m$"
	for i := m; i <= n; i++ {
		x = append(x, i)
	}
}

func forIncVarToVarMinusOneInclusive() {
	m := 0
	n := 5
	var x []int // want "Consider preallocating x with capacity n - m$"
	for i := m; i <= n-1; i++ {
		x = append(x, i)
	}
}

func forDecVarMinusOneToVarInclusive() {
	m := 0
	n := 5
	var x []int // want "Consider preallocating x with capacity n - m$"
	for i := n - 1; i >= m; i-- {
		x = append(x, i)
	}
}

func forIncVarToVarCommonOffset() {
	m := 1
	n := 5
	var x []int // want "Consider preallocating x with capacity n$"
	for i := m; i < m+n; i++ {
		x = append(x, i)
	}
}

func forIterateZeroTimes() {
	var x []int
	for i := 0; i < 0; i++ {
		x = append(x, i)
	}
}

func forIterateNegativeTimes() {
	var x []int
	for i := 1; i < 0; i++ {
		x = append(x, i)
	}
}

func forIncBackwardsCondition() {
	var x []int
	for i := 0; i > 0; i++ {
		x = append(x, i)
	}
}

func forDecBackwardsCondition() {
	var x []int
	for i := 0; i < 0; i-- {
		x = append(x, i)
	}
}

func forTypeConvert() {
	var x []uint // want "Consider preallocating x with capacity 5$"
	for i := uint(0); i < uint(5); i++ {
		x = append(x, i)
	}
}

func forNoArgMethod() {
	var buf bytes.Buffer
	var x []int // want "Consider preallocating x with capacity buf\\.Len\\(\\)$"
	for i := 0; i < buf.Len(); i++ {
		x = append(x, i)
	}
}

func forMultipleConjunctiveUpperLimits() {
	m := 7
	n := 6
	var x []int // want "Consider preallocating x with capacity min\\(m, n, 5\\)$"
	for i := 0; i < m && i < n && i < 5; i++ {
		x = append(x, i)
	}
}

func forMultipleConjunctiveUpperLimitsWithMin() {
	m := 7
	n := 6
	var x []int // want "Consider preallocating x with capacity min\\(n, 5, m\\)$"
	for i := 0; i < m && i < min(n, 5); i++ {
		x = append(x, i)
	}
}

func forMultipleDisjunctiveUpperLimits() {
	m := 3
	n := 4
	var x []int // want "Consider preallocating x with capacity max\\(m, n, 5\\)$"
	for i := 0; i < m || i < n || i < 5; i++ {
		x = append(x, i)
	}
}

func forMultipleDisjunctiveUpperLimitsWithMax() {
	m := 3
	n := 4
	var x []int // want "Consider preallocating x with capacity max\\(n, 5, m\\)$"
	for i := 0; i < m || i < max(n, 5); i++ {
		x = append(x, i)
	}
}

func forLowerLimitFunc() {
	fn := func() int { return 0 }
	var x []int // want "Consider preallocating x$"
	for i := fn(); i < 5; i++ {
		x = append(x, i)
	}
}

func forUpperLimitFunc() {
	fn := func() int { return 5 }
	var x []int // want "Consider preallocating x$"
	for i := 0; i < fn(); i++ {
		x = append(x, i)
	}
}

func forUpperSelfCap() {
	var x []int
	for i := 0; i < cap(x); i++ {
		x = append(x, i)
	}
}

func forLowerSelfCap() {
	var x []int
	for i := cap(x); i < 5; i++ {
		x = append(x, i)
	}
}

func forLinkedListTraversal() {
	type node struct {
		next *node
		id   int
	}
	var x []int
	for n := new(node); n != nil; n = n.next {
		x = append(x, n.id)
	}
}

func forIncSkipLit() {
	var x []int // want "Consider preallocating x with capacity 5$"
	for i := 0; i < 10; i += 2 {
		x = append(x, i)
	}
}

func forIncSkipLitRemainder() {
	var x []int // want "Consider preallocating x with capacity 4$"
	for i := 0; i < 10; i += 3 {
		x = append(x, i)
	}
}

func forDecSkipLit() {
	var x []int // want "Consider preallocating x with capacity 5$"
	for i := 10; i > 0; i -= 2 {
		x = append(x, i)
	}
}

func forIncSkipVar() {
	n := 10
	var x []int // want "Consider preallocating x with capacity n/2 \\+ 1$"
	for i := 0; i < n; i += 2 {
		x = append(x, i)
	}
}

func forIncSkipBinaryLit() {
	var x []int // want "Consider preallocating x with capacity 5$"
	for i := 0; i < 10; i = i + 2 {
		x = append(x, i)
	}
}

func forIncSkipBinaryLitBackwards() {
	var x []int // want "Consider preallocating x with capacity 5$"
	for i := 0; i < 10; i = 2 + i {
		x = append(x, i)
	}
}

func forMultiVarsInc() {
	var x []int // want "Consider preallocating x with capacity 20$"
	for i, j := 0, 0; i < 10; i++ {
		x = append(x, i, j)
	}
}

func forMultiVarsBinary() {
	var x []int // want "Consider preallocating x with capacity 20$"
	for i, j := 0, 0; i < 10; i, j = i+1, j+2 {
		x = append(x, i, j)
	}
}
