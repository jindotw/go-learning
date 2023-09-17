package main

type MinHeap2 []int

func (h *MinHeap2) Len() int {
	return len(*h)
}

func (h *MinHeap2) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap2) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap2) Push(x interface{}) {
	switch t := x.(type) {
	case []int:
		for _, v := range t {
			*h = append(*h, v)
		}
	case int:
		*h = append(*h, t)
	}
}

func (h *MinHeap2) Pop() interface{} {
	last := h.Len() - 1
	if last < 0 {
		return nil
	}
	val := (*h)[last]
	*h = (*h)[:last]
	return val
}
