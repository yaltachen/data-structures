package avl

import (
	"fmt"
)

type avlTree struct {
	size       int
	root       *node
	comparator Comparator
}

type node struct {
	Value     interface{}
	LeftNode  *node
	RightNode *node
	height    int
}

func getHeight(node *node) int {
	if node == nil {
		return 0
	}
	return node.height
}

func AVLTree(comparator Comparator) *avlTree {
	return &avlTree{root: nil, size: 0, comparator: comparator}
}

func (a avlTree) GetSize() int {
	return a.size
}

func GetBalanceFactor(n *node) int {
	if n == nil {
		return 0
	}
	return getHeight(n.LeftNode) - getHeight(n.RightNode)
}

func (a *avlTree) Add(value interface{}) {
	a.root = a.add(a.root, value)
}

func (a *avlTree) add(node *node, value interface{}) *node {
	if node == nil {
		a.size++
		return buildNode(value)
	}
	if a.comparator.Compare(node.Value, value) > 0 {
		// value < node.value 左子树
		node.LeftNode = a.add(node.LeftNode, value)
	} else if a.comparator.Compare(node.Value, value) < 0 {
		// value > node.value 右子树
		node.RightNode = a.add(node.RightNode, value)
	} else {
		// DO Noting
	}
	// 重新计算node的高度
	node.height = max(getHeight(node.RightNode), getHeight(node.LeftNode)) + 1

	var balance = GetBalanceFactor(node)
	if balance > 1 || balance < -1 {
		// 维护平衡
		// LL
		if balance > 1 && GetBalanceFactor(node.LeftNode) >= 0 {
			return node.rightRotate()
		}
		// RR
		if balance < -1 && GetBalanceFactor(node.LeftNode) <= 0 {
			return node.leftRotate()
		}
		// LR
		if balance > 1 && GetBalanceFactor(node.LeftNode) < 0 {
			node.LeftNode = node.LeftNode.leftRotate()
			return node.rightRotate()
		}
		// RL
		if balance < -1 && GetBalanceFactor(node.LeftNode) > 0 {
			node.RightNode = node.RightNode.rightRotate()
			return node.leftRotate()
		}
	}
	return node
}

func (a *avlTree) Remove(value interface{}) {
	a.root = a.remove(a.root, value)
}

func (a *avlTree) remove(n *node, value interface{}) *node {
	if n == nil {
		return nil
	}
	// 删除节点后的新的根节点
	var retNode *node
	if a.comparator.Compare(n.Value, value) > 0 {
		// value < n.value 向左子树找
		n.LeftNode = a.remove(n.LeftNode, value)
		retNode = n
	} else if a.comparator.Compare(n.Value, value) < 0 {
		// value > n.value 向右子树找
		n.RightNode = a.remove(n.RightNode, value)
		retNode = n
	} else {
		// 找到待删除节点
		if n.RightNode == nil {
			// 删除元素右子树为空
			var leftNode = n.LeftNode
			n.LeftNode = nil
			a.size--
			retNode = leftNode
		} else if n.LeftNode == nil {
			// 删除元素左子树为空
			var rightNode = n.RightNode
			n.RightNode = nil
			a.size--
			retNode = rightNode
		} else {
			// 左右子树均不为空
			var precursor = a.getMaximum(n.LeftNode)
			precursor.LeftNode = a.remove(n.LeftNode, precursor.Value)
			precursor.RightNode = n.RightNode
			n.LeftNode = nil
			n.RightNode = nil
			retNode = precursor
		}
	}

	if retNode == nil {
		return nil
	}

	// 重新计算高度
	retNode.height = max(
		getHeight(retNode.LeftNode),
		getHeight(retNode.RightNode),
	) + 1

	// 维护retNode平衡
	var balance = GetBalanceFactor(retNode)
	if balance > 1 || balance < -1 {
		// 维护平衡
		// LL
		if balance > 1 && GetBalanceFactor(retNode.LeftNode) >= 0 {
			return retNode.rightRotate()
		}
		// RR
		if balance < -1 && GetBalanceFactor(retNode.LeftNode) <= 0 {
			return retNode.leftRotate()
		}
		// LR
		if balance > 1 && GetBalanceFactor(retNode.LeftNode) < 0 {
			retNode.LeftNode = retNode.LeftNode.leftRotate()
			return retNode.rightRotate()
		}
		// RL
		if balance < -1 && GetBalanceFactor(retNode.LeftNode) > 0 {
			retNode.RightNode = retNode.RightNode.rightRotate()
			return retNode.leftRotate()
		}
	}
	return retNode
}

func (a *avlTree) PreOrder() {
	a.preOrder(a.root)
}

func (a *avlTree) preOrder(node *node) {
	if node == nil {
		return
	}
	fmt.Println(node.Value, GetBalanceFactor(node))
	a.preOrder(node.LeftNode)
	a.preOrder(node.RightNode)
}

func buildNode(value interface{}) *node {
	return &node{Value: value, LeftNode: nil, RightNode: nil, height: 1}
}

func (a avlTree) ISBST() bool {
	var aux = a.inOrder(a.root, make([]interface{}, 0, 0))
	for i := 0; i < len(aux)-1; i++ {
		if a.comparator.Compare((aux[i]), aux[i+1]) > 0 {
			return false
		}
	}
	return true
}

func (a avlTree) ISBalance() bool {
	return a.isBalance(a.root)
}

func (a avlTree) isBalance(node *node) bool {
	if node == nil {
		return true
	}
	if GetBalanceFactor(node) > 1 || GetBalanceFactor(node) < -1 {
		return false
	}
	return a.isBalance(node.LeftNode) && a.isBalance(node.RightNode)
}

func (a avlTree) inOrder(node *node, aux []interface{}) []interface{} {
	if node == nil {
		return aux
	}
	aux = append(aux, node.Value)
	a.inOrder(node.LeftNode, aux)
	a.inOrder(node.RightNode, aux)
	return aux
}

//		y(node)						x
//	t1		 x					y     z
//		  t2   z             t1 t2   t3 t4
//			  t3 t4
func (node *node) leftRotate() *node {
	var x = node.RightNode
	var t2 = x.LeftNode

	x.LeftNode = node
	node.RightNode = t2

	node.height = max(getHeight(node.LeftNode), getHeight(node.RightNode)) + 1
	x.height = max(getHeight(x.LeftNode), getHeight(x.RightNode)) + 1

	return x
}

//         y(node)				x
//      x    t4				z		y(node)
//    z  t3				  t1 t2	   t3 t4
//  t1 t2
func (node *node) rightRotate() *node {
	var x = node.LeftNode
	var t3 = x.RightNode

	x.RightNode = node
	node.LeftNode = t3

	// update height
	node.height = max(getHeight(node.LeftNode), getHeight(node.RightNode)) + 1
	x.height = max(getHeight(x.LeftNode), getHeight(x.RightNode)) + 1

	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (a *avlTree) getMaximum(node *node) *node {
	if node == nil {
		return nil
	}
	if node.RightNode == nil {
		return node
	}
	return a.getMaximum(node.RightNode)
}

func (a avlTree) LevelOrder() {
	var (
		lastNode        *node
		currentNextNode *node
		currentNode     *node
		queue           []*node
		levelCache      []interface{}
		level           int
	)
	level = 0
	lastNode = a.root
	currentNextNode = nil
	queue = make([]*node, 0, 0)
	levelCache = make([]interface{}, 0, 0)

	queue = append(queue, a.root)

	for len(queue) != 0 && lastNode != nil {
		currentNode = queue[0]
		queue = queue[1:]
		if currentNode != nil {
			levelCache = append(levelCache, currentNode.Value)
		} else {
			levelCache = append(levelCache, "nil")
		}
		if currentNode != nil {
			queue = append(queue, currentNode.LeftNode)
			queue = append(queue, currentNode.RightNode)
			if currentNode.RightNode != nil {
				currentNextNode = currentNode.RightNode
			} else {
				currentNextNode = currentNode.LeftNode
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
