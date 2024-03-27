// https://leetcode.com/problems/fibonacci-number

package fibonacciNumber

func fib(n int) int {
	if n <= 1 {
		return n
	}

	cache := make([]int, n+1)

	return fibCache(n, cache)
}

func fibCache(n int, cache []int) int {
	if n <= 1 {
		return n
	}

	if cache[n] != 0 {
		return cache[n]
	}

	cache[n] = fibCache(n-1, cache) + fibCache(n-2, cache)

	return cache[n]
}
