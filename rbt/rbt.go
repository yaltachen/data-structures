package rbt

import "fmt"

type rbt struct {
	root       *Node
	size       int
	comparator Comparator
}

type Node struct {
	Value     interface{}
	LeftNode  *Node
	RightNode *Node
	Color     bool
}

const (
	RED   = true
	BLACK = false
)

func RBT(comparator Comparator) *rbt {
	return &rbt{root: nil, size: 0, comparator: comparator}
}

func (r rbt) getSize() int {
	return r.size
}

func getColor(node *Node) bool {
	if node == nil {
		return BLACK
	} else {
		return node.Color
	}
}

func (rbt *rbt) Add(value interface{}) {
	rbt.root = rbt.add(rbt.root, value)
	rbt.root.Color = BLACK
}

func (rbt *rbt) add(node *Node, value interface{}) *Node {
	if node == nil {
		return &Node{Value: value, LeftNode: nil, RightNode: nil, Color: RED}
	}

	if rbt.comparator.Compare(value, node.Value) < 0 {
		node.LeftNode = rbt.add(node.LeftNode, value)
	} else if rbt.comparator.Compare(value, node.Value) > 0 {
		node.RightNode = rbt.add(node.RightNode, value)
	}

	// 维护颜色
	if getColor(node.RightNode) == RED && getColor(node.LeftNode) != RED {
		node = leftRotate(node)
	}
	if getColor(node.LeftNode) == RED && getColor(node.LeftNode.LeftNode) == RED {
		node = rightRotate(node)
	}
	if getColor(node.LeftNode) == RED && getColor(node.RightNode) == RED {
		node = colorflip(node)
	}

	return node
}

//		  y
//	  t1     x
//		   t2  t3

func leftRotate(y *Node) *Node {
	var x = y.RightNode

	y.RightNode = x.LeftNode
	x.LeftNode = y

	x.Color = y.Color
	y.Color = RED

	return x
}

//	    y
//	  x  t4
//	z  t3
//t1 t2
func rightRotate(y *Node) *Node {
	var x = y.LeftNode
	var t3 = x.RightNode
	x.LeftNode.Color = BLACK
	y.Color = BLACK
	x.Color = RED

	y.LeftNode = t3
	x.RightNode = y

	return x
}

//		    y
//	z(RED)	    x(RED)
//
func colorflip(y *Node) *Node {
	y.LeftNode.Color = BLACK
	y.RightNode.Color = BLACK
	y.Color = RED
	return y
}

func (b rbt) LevelOrder() {
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
			levelCache = append(levelCache, struct {
				value interface{}
				color bool
			}{
				value: currentNode.Value, color: currentNode.Color})
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
