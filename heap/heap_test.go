package heap

import (
	"fmt"
	"testing"
)

type IntComparator struct{}

func (IntComparator) Compare(a, b interface{}) int {
	if a.(int) > b.(int) {
		return 1
	} else if a.(int) == b.(int) {
		return 0
	} else {
		return -1
	}
}

func TestHeap(t *testing.T) {
	var heap = &Heap{arr: make([]interface{}, 0, 0), comparator: IntComparator{}}
	for i := 10; i > 0; i-- {
		heap.Add(i)
	}
	fmt.Println(heap.arr)
}
