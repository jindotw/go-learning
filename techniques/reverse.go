package main

import "fmt"

func reverseSlice() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Printf("%#v\n", arr)
}

func reverseSlice2() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - 1 - i
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Printf("%#v\n", arr)
}

func main() {
	reverseSlice2()
}
