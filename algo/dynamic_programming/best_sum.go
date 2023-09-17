package main

import "fmt"

func bestSumNoMemo(target int, nums []int) []int {
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	var result []int
	for _, num := range nums {
		ret := bestSumNoMemo(target-num, nums)
		if ret != nil {
			ret = append(ret, num)
			if result == nil || len(ret) < len(result) {
				result = ret
			}
		}
	}
	return result
}

func bestSum(target int, nums []int, memo map[int][]int) []int {
	if val, ok := memo[target]; ok {
		return val
	}
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	var result []int
	for _, num := range nums {
		ret := bestSum(target-num, nums, memo)
		if ret != nil {
			ret = append(ret, num)
			if result == nil || len(ret) < len(result) {
				result = ret
			}
		}
	}
	memo[target] = result
	return result
}

func callBestSum() {
	nums := []int{2, 5, 10, 20, 25, 50}
	target := 100
	memo := make(map[int][]int)
	fmt.Println(bestSum(target, nums, memo))
	fmt.Println(bestSumNoMemo(target, nums))
}
