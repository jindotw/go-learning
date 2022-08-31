package main

import "fmt"

func main() {
	strings := []string{"a", "b", "c"}
	for _, v := range strings {
		fmt.Println(v)
	}
	fmt.Println()
	strings = append(strings, "d")
	for _, v := range strings {
		fmt.Println(v)
	}
}
