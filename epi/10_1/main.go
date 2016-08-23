package main

import (
	"fmt"

	"github.com/hmeng-19/goutils/btree"
)

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func isHeightBalanced(root *btree.BTree) (int, int, bool) {
	if root == nil {
		return -1, -1, true
	}

	leftLeftH, leftRightH, leftBalanced := isHeightBalanced(root.Left)
	rightLeftH, rightRightH, rightBalanced := isHeightBalanced(root.Right)

	leftH := max(leftLeftH, leftRightH) + 1
	rightH := max(rightLeftH, rightRightH) + 1

	fmt.Printf("root: %v; leftH: %v; rightH: %v\n", root.Data, leftH, rightH)
	if !leftBalanced || rightBalanced == false {
		return leftH, rightH, false
	}

	diff := abs(rightH - leftH)
	if diff > 1 {
		return leftH, rightH, false
	}
	return leftH, rightH, true
}

func main() {
	root := &btree.BTree{Data: 3}
	n1 := &btree.BTree{Data: 4}
	n2 := &btree.BTree{Data: 5}
	n3 := &btree.BTree{Data: 7}
	n4 := &btree.BTree{Data: 6}
	root.Left, root.Right = n1, n2
	n1.Left, n1.Right = n3, n4
	//n5 := &btree.BTree{2, nil, nil}
	//n4.Right = n5
	root.InOrder()
	l, r, balanced := isHeightBalanced(root)
	fmt.Printf("l: %v; r: %v; balanced: %v\n", l, r, balanced)
}
