package main

type MaxHeap2 []int

func (h *MaxHeap2) Len() int {
	return len(*h)
}

func (h *MaxHeap2) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap2) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap2) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap2) Pop() interface{} {
	last := h.Len() - 1
	val := (*h)[last]
	*h = (*h)[:last]
	return val
}
