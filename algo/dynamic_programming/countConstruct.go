package main

import (
	"fmt"
	"strings"
)

func countConstructNoMemo(target string, wordBank []string) int {
	if target == "" {
		return 1
	}

	total := 0
	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			total += countConstructNoMemo(target[len(word):], wordBank)
		}
	}

	return total
}

func countConstruct(target string, wordBank []string, memo map[string]int) int {
	if val, ok := memo[target]; ok {
		return val
	}
	if target == "" {
		return 1
	}

	total := 0
	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			total += countConstruct(target[len(word):], wordBank, memo)
		}
	}

	memo[target] = total
	return total

}

func callCountConstruct() {
	wordBank := []string{"purp", "p", "ur", "le", "purpl"}
	target := "purple"
	memo := make(map[string]int)

	fmt.Println(countConstruct(target, wordBank, memo))
}
