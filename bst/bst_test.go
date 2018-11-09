package bst

import (
	"fmt"
	"testing"
)

func TestBst(t *testing.T) {
	bst := BST(IntComarator{})
	bst.Add(28)
	bst.Add(16)
	bst.Add(30)
	bst.Add(13)
	bst.Add(22)
	bst.Add(29)
	bst.Add(42)
	bst.Add(44)
	bst.Add(45)
	bst.Add(21)
	bst.Add(12)
	// bst.PreOrder()
	// bst.PreOrderNR()
	bst.LevelOrder()
	bst.Remove(28)
	fmt.Println("...")
	bst.LevelOrder()
	// bst.InOrder()
	// bst.PostOrder()
	// bst.RemoveMin()
	// bst.LevelOrder()
	// bst.RemoveMax()
	// bst.LevelOrder()
}

type IntComarator struct{}

func (IntComarator) Compare(a, b interface{}) int {
	if a.(int) > b.(int) {
		return 1
	} else if a.(int) < b.(int) {
		return -1
	} else {
		return 0
	}
}
