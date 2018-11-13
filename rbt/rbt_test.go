package rbt

import (
	"testing"
)

func TestRBT(t *testing.T) {
	var rbt = RBT(IntComparator{})
	for i := 0; i < 10; i++ {
		rbt.Add(i)
	}
	rbt.LevelOrder()
}

type IntComparator struct{}

func (IntComparator) Compare(a, b interface{}) int {
	if a.(int) > b.(int) {
		return 1
	} else if a.(int) == b.(int) {
		return 0
	}
	return -1
}
