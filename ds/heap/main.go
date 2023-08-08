package main

import "fmt"

func main() {
	arr := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	maxHeap := &MaxHeap{}
	*maxHeap = arr
	maxHeap.BuildHeap(arr)
	fmt.Println(*maxHeap)

	fmt.Println(maxHeap.Pop())
	fmt.Println(*maxHeap)
	fmt.Println(maxHeap.Pop())
	fmt.Println(*maxHeap)
	fmt.Println(maxHeap.Pop())
}
