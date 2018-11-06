package segmentTree

import (
	"bytes"
)

type segmentTree struct {
	tree    []interface{}
	data    []interface{}
	merger  Merger
	printer Printer
}

func (st *segmentTree) GetSize() int {
	return len(st.data)
}

// 创建segmenttree
func SegmentTree(arr []interface{}, merger Merger, printer Printer) *segmentTree {
	var (
		i    int
		tree []interface{}
		data []interface{}
		st   *segmentTree
	)
	for i = 0; i < len(arr); i++ {
		data = append(data, arr[i])
	}
	tree = make([]interface{}, 4*len(arr), 4*len(arr))
	st = &segmentTree{tree: tree, data: data, merger: merger, printer: printer}
	// 创建树结构
	st.tree[0] = st.buildSegmentTree(0, 0, len(data)-1)
	return st
}

func (st *segmentTree) Get(index int) interface{} {
	if index < 0 || index >= len(st.data) {
		panic("index out of range")
	}
	return st.tree[index]
}

func (st *segmentTree) buildSegmentTree(rootIndex, left, right int) interface{} {
	if left == right {
		st.tree[rootIndex] = st.data[left]
		return st.tree[rootIndex]
	}

	var mid = left + (right-left)/2
	var leftChild = getLeftChild(rootIndex)
	var rightChild = getRightChild(rootIndex)

	st.tree[leftChild] = st.buildSegmentTree(leftChild, left, mid)
	st.tree[rightChild] = st.buildSegmentTree(rightChild, mid+1, right)

	return st.merger.Merege(st.tree[leftChild], st.tree[rightChild])
}

func (st *segmentTree) String() string {
	var (
		ele    interface{}
		buffer bytes.Buffer
	)
	buffer.WriteString("[")
	for _, ele = range st.tree {
		if ele != nil {
			buffer.WriteString(st.printer.ToString(ele))
		} else {
			buffer.WriteString("null")
		}
		buffer.WriteString(",")
	}
	buffer.WriteString("]")
	return buffer.String()
}

func (st *segmentTree) Query(lQuery, rQuery int) (res interface{}) {
	if lQuery > rQuery || lQuery < 0 || rQuery < 0 || lQuery >= len(st.data) {
		panic("index illegal")
	}
	return st.query(0, 0, len(st.data)-1, lQuery, rQuery)
}

func (st *segmentTree) query(rootIndex, lRange, rRange, lQuery, rQuery int) interface{} {
	// 找到指定范围
	if lRange == lQuery && rRange == rQuery {
		return st.tree[rootIndex]
	}

	var mid = lRange + (rRange-lRange)/2
	var leftChild = getLeftChild(rootIndex)
	var rightChild = getRightChild(rootIndex)

	// 范围在rootIndex左子树:
	if rQuery <= mid {
		return st.query(leftChild, lRange, mid, lQuery, rQuery)
	}
	// 范围在rootIndex右子树:
	if lQuery >= mid+1 {
		return st.query(rightChild, mid+1, rRange, lQuery, rQuery)
	}

	// 范围跨越
	return st.merger.Merege(
		st.query(leftChild, lRange, mid, lQuery, mid),      // 左边的
		st.query(rightChild, mid+1, rRange, mid+1, rQuery), // 右边的
	)
}

func (st *segmentTree) Set(index int, e interface{}) {
	if index < 0 || index >= len(st.data) {
		panic("index out of range")
	}
	st.data[index] = e
	st.set(0, 0, len(st.data)-1, index, e)
}

func (st *segmentTree) set(treeIndex, leftIndex, rightIndex, index int, e interface{}) {
	if leftIndex == rightIndex {
		// 找到要更新节点
		st.tree[treeIndex] = e
		return
	}
	var mid = leftIndex + (rightIndex-leftIndex)/2
	var leftChild = getLeftChild(treeIndex)
	var rigtChild = getRightChild(treeIndex)

	if index > mid+1 {
		st.set(rigtChild, mid+1, rightIndex, index, e)
	} else {
		st.set(leftChild, leftIndex, mid, index, e)
	}
	st.tree[treeIndex] = st.merger.Merege(
		st.tree[leftChild],
		st.tree[rigtChild],
	)
}

func getLeftChild(rootIndex int) int {
	return rootIndex*2 + 1
}

func getRightChild(rootIndex int) int {
	return rootIndex*2 + 2
}
