package main

import (
	"container/heap"
	"fmt"
)

func main() {
	minHeap := &MinHeap2{}
	a := []int{1, 4, 5}
	b := []int{1, 3, 4}
	c := []int{2, 6}
	for _, num := range a {
		heap.Push(minHeap, num)
	}
	for _, num := range b {
		heap.Push(minHeap, num)
	}
	for _, num := range c {
		heap.Push(minHeap, num)
	}
	for i, v := range *minHeap {
		fmt.Printf("[%d] %d\n", i, v)
	}
	for minHeap.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(minHeap))
	}
	fmt.Println()
}
