package bst

import (
	"fmt"
)

type bst struct {
	root       *Node
	size       int
	comparator Comparator
}

type Node struct {
	Value      interface{}
	LeftChild  *Node
	RightChild *Node
}

func (b bst) GetSize() int {
	return b.size
}

func (b bst) IsEmpty() bool {
	if b.size <= 0 {
		return true
	}
	return false
}

func BST(comparator Comparator) *bst {
	return &bst{root: nil, size: 0, comparator: comparator}
}

// 向bst中插入元素
func (b *bst) Add(value interface{}) {
	b.root = b.add(b.root, value)
}

// 向以rootNode为根节点的子树中，插入元素
// 并返回插入元素后，bst的根节点
func (b *bst) add(rootNode *Node, value interface{}) *Node {
	// 递归出口
	// rootNode为空，说明到了叶子节点，此时新建节点
	if rootNode == nil {
		b.size++
		return &Node{Value: value, LeftChild: nil, RightChild: nil}
	}

	if b.comparator.Compare(rootNode.Value, value) > 0 {
		// value < 当前节点的值，向rootNode的左孩子插入
		rootNode.LeftChild = b.add(rootNode.LeftChild, value)
	} else if b.comparator.Compare(rootNode.Value, value) < 0 {
		// value > 当前节点的值，向rootNode的右孩子插入
		rootNode.RightChild = b.add(rootNode.RightChild, value)
	}
	// value = 当前节点的值，直接返回
	return rootNode
}

// 查询元素
func (b bst) IsContains(value interface{}) bool {
	return b.isContains(b.root, value)
}

// 查询以rootNode为根的树中，是否包含元素
func (b bst) isContains(rootNode *Node, value interface{}) bool {
	// 递归出口，rootNode为空，表示没有找到
	if rootNode == nil {
		return false
	}

	// value = rootNode.value return true
	if b.comparator.Compare(rootNode.Value, value) == 0 {
		return true
	} else if b.comparator.Compare(rootNode.Value, value) > 0 {
		// rootNode.value > value 向左子树找
		return b.isContains(rootNode.LeftChild, value)
	} else {
		// rootNode.value < value 向右子树找
		return b.isContains(rootNode.RightChild, value)
	}
}

// 前序遍历
func (b bst) PreOrder() {
	b.preOrder(b.root)
}

func (b bst) preOrder(rootNode *Node) {
	// 递归出口
	if rootNode == nil {
		return
	}
	fmt.Printf("%v\r\n", rootNode.Value)
	b.preOrder(rootNode.LeftChild)
	b.preOrder(rootNode.RightChild)
}

// 前序遍历（非递归）
func (b bst) PreOrderNR() {

	// 栈
	var stack = make([]*Node, 0, 0)
	var node *Node

	stack = append(stack, b.root)

	for len(stack) != 0 {
		// 出栈
		node = stack[len(stack)-1]
		fmt.Println(node.Value)
		stack = stack[0 : len(stack)-1]
		// 压右孩子
		if node.RightChild != nil {
			stack = append(stack, node.RightChild)
		}
		// 压左孩子
		if node.LeftChild != nil {
			stack = append(stack, node.LeftChild)
		}
	}

}

// 中序遍历
func (b bst) InOrder() {
	b.inOrder(b.root)
}

func (b bst) inOrder(rootNode *Node) {
	// 递归出口
	if rootNode == nil {
		return
	}
	b.inOrder(rootNode.LeftChild)
	fmt.Printf("%v\r\n", rootNode.Value)
	b.inOrder(rootNode.RightChild)
}

// 后序遍历
func (b bst) PostOrder() {
	b.postOrder(b.root)
}

func (b bst) postOrder(rootNode *Node) {
	// 递归出口
	if rootNode == nil {
		return
	}
	b.postOrder(rootNode.LeftChild)
	b.postOrder(rootNode.RightChild)
	fmt.Printf("%v\r\n", rootNode.Value)
}

// 层序遍历
func (b bst) LevelOrder() {
	var (
		lastNode        *Node
		currentNextNode *Node
		currentNode     *Node
		queue           []*Node
		levelCache      []interface{}
		level           int
	)
	level = 0
	lastNode = b.root
	currentNextNode = nil
	queue = make([]*Node, 0, 0)
	levelCache = make([]interface{}, 0, 0)

	queue = append(queue, b.root)

	for len(queue) != 0 && lastNode != nil {
		currentNode = queue[0]
		queue = queue[1:]
		if currentNode != nil {
			levelCache = append(levelCache, currentNode.Value)
		} else {
			levelCache = append(levelCache, "nil")
		}
		if currentNode != nil {
			queue = append(queue, currentNode.LeftChild)
			queue = append(queue, currentNode.RightChild)
			if currentNode.RightChild != nil {
				currentNextNode = currentNode.RightChild
			} else {
				currentNextNode = currentNode.LeftChild
			}
		}
		if lastNode == currentNode {
			lastNode = currentNextNode
			for getLevelCount(level) != len(levelCache) {
				levelCache = append(levelCache, "nil")
			}
			fmt.Printf("第%d层: %v\r\n", level, levelCache)
			levelCache = levelCache[0:0]
			level++
		}
	}
}

func getLevelCount(level int) int {
	var count = 1
	for i := 0; i < level; i++ {
		count = count * 2
	}
	return count
}

// 删除BST中值最小的节点，并返回删除节点的值
func (b *bst) RemoveMin() interface{} {
	var minNode = b.getMinimum(b.root)
	b.root = b.removeMin(b.root)
	return minNode.Value
}

// 删除以rootNode为根节点的BST中值最小的节点
// 并返回删除节点后的新的根节点
func (b *bst) removeMin(rootNode *Node) *Node {
	// 递归出口
	if rootNode.LeftChild == nil {
		var rightChild = rootNode.RightChild
		rootNode.RightChild = nil
		b.size--
		return rightChild
	}

	rootNode.LeftChild = b.removeMin(rootNode.LeftChild)
	return rootNode
}

// 删除BST中值最大的节点，并返回删除节点的值
func (b *bst) RemoveMax() interface{} {
	var maxNode = b.getMaximum(b.root)
	b.root = b.removeMax(b.root)
	return maxNode.Value
}

// 删除已rootNode为根节点的BST中值最大的节点
// 并返回删除节点后的新的根节点
func (b *bst) removeMax(rootNode *Node) *Node {
	if rootNode.RightChild == nil {
		var leftChild = rootNode.LeftChild
		rootNode.LeftChild = nil
		b.size--
		return leftChild
	}
	rootNode.RightChild = b.removeMax(rootNode.RightChild)
	return rootNode
}

// 返回BST中值最小的节点的值
func (b bst) GetMinimum() interface{} {
	if b.IsEmpty() {
		return nil
	} else {
		return b.getMinimum(b.root).Value
	}
}

// 返回以rootNode为根的BST中值最小的节点
func (b bst) getMinimum(rootNode *Node) *Node {
	if rootNode.LeftChild == nil {
		return rootNode
	}
	return b.getMinimum(rootNode.LeftChild)
}

// 返回BST中值最大的节点的值
func (b bst) GetMaximum() interface{} {
	if b.IsEmpty() {
		return nil
	} else {
		return b.getMaximum(b.root).Value
	}
}

// 返回以rootNode为根的BST中值最大的节点
func (b bst) getMaximum(rootNode *Node) *Node {
	if rootNode.RightChild == nil {
		return rootNode
	}
	return b.getMaximum(rootNode.RightChild)
}

// 删除值为value的节点
func (b *bst) Remove(value interface{}) {
	b.root = b.remove(b.root, value)
}

// 删除以rootNode为根的BST中的值为value的节点
// 并返回删除节点后新的rootNode
func (b *bst) remove(rootNode *Node, value interface{}) *Node {
	// 不包含value
	if rootNode == nil {
		return rootNode
	}

	if b.comparator.Compare(rootNode.Value, value) > 0 {
		// 当前节点.value > target value 向左子树找
		rootNode.LeftChild = b.remove(rootNode.LeftChild, value)
	} else if b.comparator.Compare(rootNode.Value, value) < 0 {
		// 当前节点.value < target value 向右子树找
		rootNode.RightChild = b.remove(rootNode.RightChild, value)
	} else {
		// 待删除的节点
		b.size--
		if rootNode.LeftChild == nil {
			// 如果左子树是空
			var rightChild = rootNode.RightChild
			rootNode.RightChild = nil
			return rightChild
		}
		if rootNode.RightChild == nil {
			// 如果右子树是空
			var leftChild = rootNode.LeftChild
			rootNode.LeftChild = nil
			return leftChild
		}
		// 均不为空
		// 找到targetNode的前驱
		var precursorNode = b.getMaximum(rootNode.LeftChild)

		precursorNode.LeftChild = b.removeMax(rootNode.LeftChild)
		precursorNode.RightChild = rootNode.RightChild

		rootNode.LeftChild = nil
		rootNode.RightChild = nil

		return precursorNode

	}
	return rootNode
}
