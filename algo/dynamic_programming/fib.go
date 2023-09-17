package main

import "fmt"

func fibNoMemo(n int) int {
	if n <= 2 {
		return 1
	}

	return fibNoMemo(n-1) + fibNoMemo(n-2)
}

func fib(n int, memo []int) int {
	if memo[n] != 0 {
		return memo[n]
	}
	if n <= 2 {
		return 1
	}
	memo[n] = fib(n-1, memo) + fib(n-2, memo)
	return memo[n]
}

func fibIter(n int) int {
	if n <= 2 {
		return 1
	}
	memo := []int{1, 1}
	for i := 3; i <= n; i++ {
		memo[0], memo[1] = memo[1], memo[0]+memo[1]
	}

	return memo[1]
}

func callFib() {
	n := 10
	memo := make([]int, n+1)
	fmt.Printf("Fibonacci of %d is %d\n", n, fibIter(n))
	fmt.Printf("Fibonacci of %d is %d\n", n, fib(n, memo))
}

// 1 1 2 3 5 8 13
