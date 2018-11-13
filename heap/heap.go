package heap

type Heap struct {
	arr        []interface{}
	comparator Comparator
}

func (h *Heap) GetSize() int {
	return len(h.arr)
}

func getParentIndex(index int) int {
	var pIndex = (index - 1) / 2
	if pIndex < 0 {
		pIndex = 0
	}
	return pIndex
}

func getLeftIndex(index int) int {
	return index*2 + 1
}

func getRightIndex(index int) int {
	return index*2 + 2
}

func shiftDown() {

}

func (h *Heap) shiftUp(index int) {
	var nodeIndex = index
	var parentIndex = getParentIndex(len(h.arr) - 1)

	for parentIndex >= 0 {
		// value < parent.value swap
		if h.comparator.Compare(h.arr[nodeIndex], h.arr[parentIndex]) < 0 {
			h.arr[nodeIndex], h.arr[parentIndex] =
				h.arr[parentIndex], h.arr[nodeIndex]
			nodeIndex = parentIndex
			parentIndex = getParentIndex(nodeIndex)
		} else {
			return
		}
	}
}

func (h *Heap) Add(value interface{}) {
	h.arr = append(h.arr, value)
	h.shiftUp(len(h.arr) - 1)
}
