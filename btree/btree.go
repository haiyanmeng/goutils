/*Package btree implements binary trees.*/
package btree

import (
	"fmt"
)

// A BTree represents a binary tree.
type BTree struct {
	Data        int
	Left, Right *BTree
}

// InOrder traverses the binary tree in-order.
func (root *BTree) InOrder() {
	if root == nil {
		return
	}
	InOrderTraverse(root.Left)
	fmt.Println(root.Data) // root.Data == (*root).Data
	InOrderTraverse(root.Right)
}

// IsBST checks whether a binary tree is BST.
func (root *BTree) IsBST(last *int, first *bool) bool {
	if root == nil {
		return true
	}

	if !root.Left.IsBST(last, first) {
		return false
	}

	if !(*first) && root.Data < *last {
		fmt.Printf("first = %v; current = %v; last = %v\n", *first, root.Data, *last)
		return false
	}
	*last = root.Data
	fmt.Printf("last is %v, first is %v\n", *last, *first)
	if *first {
		*first = false
	}

	return root.Right.IsBST(last, first)
}

// InOrderTraverse traverses the binary tree in-order.
func InOrderTraverse(root *BTree) {
	if root == nil {
		return
	}
	InOrderTraverse(root.Left)
	fmt.Println(root.Data) // root.Data == (*root).Data
	InOrderTraverse(root.Right)
}

// TraverseResult represents the traverse result.
type TraverseResult []int

// InOrderTraverseSlice traverses a binary tree in-order and
// collects the result into a TraverseResult.
func InOrderTraverseSlice(root *BTree, r *TraverseResult) {
	if root == nil {
		return
	}
	InOrderTraverseSlice(root.Left, r)
	*r = append(*r, root.Data)
	InOrderTraverseSlice(root.Right, r)
}

// IsIncreasingSlice checks whether a TraverseResult is increasing.
func (r TraverseResult) IsIncreasingSlice() bool {
	n := len(r)
	for i := 0; i < (n - 1); i++ {
		if r[i] > r[i+1] {
			return false
		}
	}
	return true
}
