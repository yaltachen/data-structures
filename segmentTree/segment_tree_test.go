package segmentTree

import (
	"fmt"
	"strconv"
	"testing"
)

type mMerger struct {
}

func (mMerger) Merege(left, right interface{}) interface{} {
	return left.(int) + right.(int)
}

type mPrinter struct {
}

func (mPrinter) ToString(ele interface{}) string {
	return strconv.Itoa(ele.(int))
}

var arr = []interface{}{-2, 0, 3, -5, 2, -1}

func TestSegmentBuilder(t *testing.T) {
	fmt.Println(SegmentTree(arr, mMerger{}, mPrinter{}))
}

func TestSegmentQuery(t *testing.T) {
	var cases = []struct {
		no     int
		left   int
		right  int
		result int
	}{
		{no: 1, left: 0, right: 2, result: 1},
		{no: 2, left: 2, right: 5, result: -1},
		{no: 3, left: 0, right: 5, result: -3},
	}
	var segmentTree = SegmentTree(arr, mMerger{}, mPrinter{})
	for _, c := range cases {
		realResult := segmentTree.Query(c.left, c.right)
		if c.result != realResult.(int) {
			t.Errorf("No: %d test case failed. Should get %d, but got %v", c.no, c.result, realResult)
		}
	}
}
