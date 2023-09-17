package main

import (
	"fmt"
	"strings"
)

func canConstructNoMemo(target string, wordBank []string) bool {
	if target == "" {
		return true
	}
	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			if canConstructNoMemo(target[len(word):], wordBank) {
				return true
			}
		}
	}

	return false
}

func canConstruct(target string, wordBank []string, memo map[string]bool) bool {
	if val, ok := memo[target]; ok {
		return val
	}
	if target == "" {
		return true
	}
	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			if canConstruct(target[len(word):], wordBank, memo) {
				memo[target] = true
				return true
			}
		}
	}

	memo[target] = false
	return false
}

func callCanConstruct() {
	target := "abcdef"
	wordBank := []string{"ab", "abc", "cd", "def", "abcd"}
	memo := make(map[string]bool)
	fmt.Println(canConstruct(target, wordBank, memo))
}
