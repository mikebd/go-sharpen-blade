// https://leetcode.com/problems/fibonacci-number

package fibonacciNumber

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
