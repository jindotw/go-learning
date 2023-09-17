package main

import "fmt"

func partition(t []int) int {
	lft := 0
	rgt := len(t) - 1

	for lft < rgt {
		if t[lft+1] <= t[lft] {
			t[lft], t[lft+1] = t[lft+1], t[lft]
			lft++
		} else if t[rgt] > t[lft] {
			rgt--
		} else {
			t[lft+1], t[rgt] = t[rgt], t[lft+1]
		}
	}
	fmt.Println(lft, t)

	return lft
}

func qs(t []int, k int) int {
	p := partition(t)
	if p == k {
		return t[k]
	} else if k < p {
		return qs(t[:p], k)
	} else {
		return qs(t[p+1:], k-(p+1))
	}
}

func main() {
	// 1 2 3 6 7 8 9
	nums := []int{7, 2, 1, 9, 3, 6, 8}
	k := 3
	fmt.Println(qs(nums, k))
}
