package main

type MaxHeap []int

func (h *MaxHeap) BuildHeap(arr []int) {
	lastNoneLeafIdx := (len(arr) - 2) / 2
	lastIdx := len(arr) - 1

	for currIdx := lastNoneLeafIdx; currIdx >= 0; currIdx-- {
		h.siftDown(currIdx, lastIdx)
	}
}

func (h *MaxHeap) Push(val int) {
	*h = append(*h, val)
	h.siftUp()
}

func (h *MaxHeap) Pop() int {
	last := len(*h) - 1
	// swap the first element with the last element
	h.swap(0, last)
	val := (*h)[last]
	// pop the last element
	*h = (*h)[:last]
	// call siftDown to update heap ordering
	h.siftDown(0, last-1)

	return val
}

func (h *MaxHeap) siftDown(currIdx, endIdx int) {
	lftChild := currIdx*2 + 1

	for lftChild <= endIdx {
		rgtChild := lftChild + 1
		if rgtChild > endIdx {
			rgtChild = -1
		}

		// get the larger child node to swap
		idxToSwap := lftChild
		if rgtChild != -1 && (*h)[lftChild] < (*h)[rgtChild] {
			idxToSwap = rgtChild
		}

		if (*h)[currIdx] > (*h)[idxToSwap] {
			break
		}

		h.swap(currIdx, idxToSwap)
		currIdx = idxToSwap
		lftChild = currIdx*2 + 1
	}
}

func (h *MaxHeap) siftUp() {
	currIdx := len(*h) - 1
	parentIdx := (currIdx - 1) / 2

	for currIdx > 0 && (*h)[currIdx] > (*h)[parentIdx] {
		h.swap(currIdx, parentIdx)
		currIdx = parentIdx
		parentIdx = (currIdx - 1) / 2
	}

}

func (h *MaxHeap) swap(a, b int) {
	(*h)[a], (*h)[b] = (*h)[b], (*h)[a]
}
