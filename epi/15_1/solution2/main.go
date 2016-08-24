package main

import (
	"fmt"

	"github.com/hmeng-19/goutils/btree"
)

func isBST(b *btree.BTree, last *int) bool {
	if b == nil {
		return true
	}

	if isBST(b.Left, last) == false {
		return false
	}

	fmt.Printf("b.Data = %v; last = %v\n", b.Data, *last)
	if b.Data < *last {
		return false
	}
	*last = b.Data

	return isBST(b.Right, last)
}

func leftMost(b *btree.BTree) (int, error) {
	if b == nil {
		return 0, fmt.Errorf("the BTree is empty")
	}
	var result int
	for b != nil {
		result = b.Data
		b = b.Left
	}
	return result, nil
}

func checkBST(b *btree.BTree) bool {
	if b == nil {
		return true
	}

	last, _ := leftMost(b)
	fmt.Printf("leftMost: %v\n", last)
	return isBST(b, &last)
}

func isBST1(root *btree.BTree) (int, int, bool) {
	if root == nil {
		return 0, 0, true
	}

	if root.Left == nil && root.Right == nil {
		return root.Data, root.Data, true
	}

	var min, max int
	if root.Left != nil {
		leftMin, leftMax, leftB := isBST1(root.Left)
		if leftMax > root.Data || leftB == false {
			return 0, 0, false
		}
		min = leftMin
	} else {
		min = root.Data
	}

	if root.Right != nil {
		rightMin, rightMax, rightB := isBST1(root.Right)
		if rightMin < root.Data || rightB == false {
			return 0, 0, false
		}
		max = rightMax
	} else {
		max = root.Data
	}
	return min, max, true
}

func main() {
	n1 := &btree.BTree{3, nil, nil}
	n2 := &btree.BTree{2, nil, nil}
	n3 := &btree.BTree{4, nil, nil}
	n4 := &btree.BTree{1, nil, nil}
	n5 := &btree.BTree{3, nil, nil}
	n1.Left, n1.Right = n2, n3
	n2.Left, n2.Right = n4, n5
	fmt.Println(checkBST(n1))
	_, _, x := isBST1(n1)
	fmt.Println(x)
}
