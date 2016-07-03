package main

import (
	"fmt"
	"github.com/hmeng-19/GoUtils/btree"
)

type limits struct {
	hasLow   bool
	low      int
	hasUpper bool
	upper    int
}

// IsBST checks whether a binary tree is a BST.
func IsBST(root *btree.BTree, l limits) bool {
	if root == nil {
		return true
	}

	if (l.hasLow && root.Data < l.low) || // || must be put into the end of the line, can not be put into the start of the next line
		(l.hasUpper && root.Data > l.upper) {
		return false
	}

	return IsBST(root.Left, limits{l.hasLow, l.low, true, root.Data}) &&
		IsBST(root.Right, limits{
			hasLow:   true,
			low:      root.Data,
			hasUpper: l.hasUpper,
			upper:    l.upper, // the ending , is necessary
		})

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
	n9 := btree.BTree{50, nil, nil}
	n1.Left, n1.Right = &n2, &n3
	n2.Left, n2.Right = &n4, &n5
	n3.Left, n3.Right = &n6, &n7
	n5.Left, n5.Right = &n8, &n9
	n1.InOrder()
	btree.InOrderTraverse(&n1)
	r := btree.TraverseResult{}
	btree.InOrderTraverseSlice(&n1, &r)
	fmt.Println(r)
	fmt.Printf("r is inscreasing? %v\n", r.IsIncreasingSlice())
	l := limits{
		hasLow:   false,
		hasUpper: false,
	}
	fmt.Printf("n1 is BST? %v\n", IsBST(&n1, l))
	v := n1.Data
	first := true
	fmt.Printf("n1.IsBST() = %v\n", n1.IsBST(&v, &first))
}
