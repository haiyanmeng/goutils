package main

import (
	"fmt"
	"github.com/hmeng-19/GoUtils/btree"
)

type searchRange struct {
	Low, High int
}

// IsBST checks whether a binary tree is a BST.
func RangeSearch(root *btree.BTree, r searchRange, result *[]int) {
	if root == nil {
		return
	}

	if root.Data < r.Low {
		RangeSearch(root.Right, r, result)
		return
	}

	if root.Data > r.High {
		RangeSearch(root.Left, r, result)
		return
	}

	RangeSearch(root.Left, searchRange{r.Low, root.Data}, result)
	*result = append(*result, root.Data)
	RangeSearch(root.Right, searchRange{root.Data, r.High}, result)
}

func main() {
	n1 := btree.BTree{5, nil, nil}
	n2 := btree.BTree{3, nil, nil}
	n3 := btree.BTree{14, nil, nil}
	n4 := btree.BTree{2, nil, nil}
	n5 := btree.BTree{4, nil, nil}
	n6 := btree.BTree{10, nil, nil}
	n7 := btree.BTree{19, nil, nil}
	n8 := btree.BTree{3, nil, nil}
	n9 := btree.BTree{5, nil, nil}
	n1.Left, n1.Right = &n2, &n3
	n2.Left, n2.Right = &n4, &n5
	n3.Left, n3.Right = &n6, &n7
	n5.Left, n5.Right = &n8, &n9

	v := n1.Data
	first := true
	if !n1.IsBST(&v, &first) {
		fmt.Println("The btree is not a BST!")
		return
	}

	result := []int{}
	RangeSearch(&n1, searchRange{3, 15}, &result)
	fmt.Println(result)
}
