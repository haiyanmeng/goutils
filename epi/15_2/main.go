package main

import (
	"fmt"
	"os"
)

type Node struct {
	x           int
	left, right *Node
}

func findNext(root *Node, target int) (int, error) {
	found := false
	result := 0

	if root == nil {
		return result, fmt.Errorf("the BST is empty")
	}

	cur := root

	for cur != nil {
		if cur.x > target {
			found = true
			result = cur.x
			cur = cur.left
		} else {
			cur = cur.right
		}

	}

	if !found {
		return result, fmt.Errorf("there is no key whose value > %v\n", target)
	}
	return result, nil
}

func main() {
	n1 := &Node{19, nil, nil}
	n2 := &Node{7, nil, nil}
	n3 := &Node{43, nil, nil}
	n4 := &Node{3, nil, nil}
	n5 := &Node{11, nil, nil}
	n6 := &Node{23, nil, nil}
	n7 := &Node{47, nil, nil}
	n8 := &Node{2, nil, nil}
	n9 := &Node{5, nil, nil}
	n10 := &Node{17, nil, nil}
	n11 := &Node{37, nil, nil}
	n12 := &Node{53, nil, nil}
	n13 := &Node{13, nil, nil}
	n14 := &Node{29, nil, nil}
	n15 := &Node{41, nil, nil}
	n16 := &Node{31, nil, nil}
	n1.left, n1.right = n2, n3
	n2.left, n2.right = n4, n5
	n3.left, n3.right = n6, n7
	n4.left, n4.right = n8, n9
	n5.right = n10
	n6.right = n11
	n7.right = n12
	n10.left = n13
	n11.left, n11.right = n14, n15
	n14.right = n16
	result, err := findNext(n1, 12)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(result)
}
