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

	// 当m为1时，其划分必然全为1
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

	return 0
}

func TestDivideInteger() {
	fmt.Print("6的划分数:", DivideInteger(6, 6))
}
