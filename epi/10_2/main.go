package main

import (
	"fmt"

	"github.com/hmeng-19/goutils/stack"
)

type Tree struct {
	data        int
	left, right *Tree
}

func compareTwoTrees(roota, rootb *Tree) bool {
	if roota == nil && rootb == nil {
		return true
	}

	if roota == nil || rootb == nil {
		return false
	}

	if roota.data != rootb.data {
		return false
	}

	return compareTwoTrees(roota.left, rootb.right) &&
		compareTwoTrees(roota.right, rootb.left)
}

func IsSymmetricTree(root *Tree) bool {
	if root == nil {
		return true
	}
	return compareTwoTrees(root.left, root.right)
}

func IsSymmetricTree1(root *Tree) bool {
	if root == nil || (root.left == nil && root.right == nil) {
		return true
	}

	if root.left == nil || root.right == nil {
		return false
	}

	s1, s2 := stack.NewGenericStack(), stack.NewGenericStack()
	s1.Push(interface{}(root.left))
	s2.Push(interface{}(root.right))

	for {
		if s1.IsEmpty() && s2.IsEmpty() {
			return true
		}

		if s1.IsEmpty() || s2.IsEmpty() {
			return false
		}

		v1, _ := s1.Pop()
		v2, _ := s2.Pop()
		node1, node2 := v1.(*Tree), v2.(*Tree)
		if node1.data != node2.data {
			return false
		}

		if node1.right != nil {
			s1.Push(interface{}(node1.right))
		}
		if node1.left != nil {
			s1.Push(interface{}(node1.left))
		}

		if node2.left != nil {
			s2.Push(interface{}(node2.left))
		}
		if node2.right != nil {
			s2.Push(interface{}(node2.right))
		}
	}
}

func main() {
	t1 := &Tree{5, nil, nil}
	t2 := &Tree{4, nil, nil}
	t3 := &Tree{4, nil, nil}
	t4 := &Tree{8, nil, nil}
	t5 := &Tree{7, nil, nil}
	t6 := &Tree{7, nil, nil}
	t7 := &Tree{8, nil, nil}
	t8 := &Tree{6, nil, nil}
	t9 := &Tree{4, nil, nil}
	t10 := &Tree{4, nil, nil}
	t11 := &Tree{6, nil, nil}
	t1.left, t1.right = t2, t3
	t2.left, t2.right = t4, t5
	t3.left, t3.right = t6, t7
	t4.left, t4.right = t8, t9
	t7.left, t7.right = t10, t11
	fmt.Println(IsSymmetricTree(t1))
	fmt.Println(IsSymmetricTree1(t1))
}
