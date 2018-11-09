package avl

import (
	"fmt"
	"testing"
)

func TestAVL(t *testing.T) {
	avlTree := AVLTree(IntComarator{})
	avlTree.Add(1)
	avlTree.Add(2)
	avlTree.Add(3)
	avlTree.Add(4)
	avlTree.Add(5)
	avlTree.Add(6)
	avlTree.Add(7)
	avlTree.Add(8)
	avlTree.Add(9)

	// avlTree.LevelOrder()
	// fmt.Println(avlTree.ISBalance(), avlTree.ISBST(), avlTree.GetSize())

	avlTree.Remove(4)
	avlTree.Remove(3)
	avlTree.LevelOrder()
	fmt.Println(avlTree.ISBalance(), avlTree.ISBST(), avlTree.GetSize())
	// avlTree.Remove(28)
	// fmt.Println(avlTree.ISBalance(), avlTree.ISBST(), avlTree.GetSize())

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
