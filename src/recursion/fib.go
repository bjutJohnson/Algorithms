package recursion

func Fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}

func FibNotRec(n int) int {
	fibSlice := make([]int, 0)

	fibSlice = append(fibSlice, 1)
	fibSlice = append(fibSlice, 1)

	iLen := len(fibSlice)

	if n >= iLen {
		for i := iLen; i <= n; i++ {
			fibSlice = append(fibSlice, fibSlice[i-1]+fibSlice[i-2])
		}
	}

	return fibSlice[n]
}
