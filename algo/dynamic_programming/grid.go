package main

import "fmt"

func gridNoMemo(m, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	return gridNoMemo(m-1, n) + gridNoMemo(m, n-1)
}

func grid(m, n int, memo map[string]int) int {
	key := fmt.Sprintf("%d,%d", m, n)
	if val, ok := memo[key]; ok {
		return val
	}
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	memo[key] = grid(m-1, n, memo) + grid(m, n-1, memo)
	return memo[key]
}

func gridIter(m, n int) int {
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		memo[i][0] = 1
	}
	for j := 0; j < n; j++ {
		memo[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			memo[i][j] = memo[i-1][j] + memo[i][j-1]
		}
	}

	return memo[m-1][n-1]
}

func callGrid() {
	m, n := 10, 10
	memo := make(map[string]int)
	fmt.Printf("A grid of %dx%d has %d ways\n", m, n, grid(m, n, memo))
	fmt.Printf("A grid of %dx%d has %d ways\n", m, n, gridIter(m, n))
}
