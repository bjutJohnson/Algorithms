package recursion

import (
	"fmt"
)

// 求整数的不同划分数， 要求n > 0
// n: 待划分的整数，m: 划分中最大的整数
func DivideInteger(n, m int) int {
	if n <= 0 {
		return 0
	}

	if m == 1 {
		return 1
	}

	if n < m {
		return DivideInteger(n, n)
	}

	if n == m {
		return 1 + DivideInteger(n, n - 1)
	}

	if n > m && n > 1 && m > 1{
		return DivideInteger(n, m - 1) + DivideInteger(n - m, m)
	}

	// 当前面的条件都不成立时，程序能执行到这里，必然是m为1的时候，显然，对于任意的n，必然只有一种分法，即都为1
	return 0
}

func TestDivideInteger() {
	fmt.Print("6的划分数:", DivideInteger(6, 6))
}
