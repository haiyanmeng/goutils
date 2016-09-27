package main

import (
	"fmt"

	"github.com/hmeng-19/goutils/stack"
)

type Node struct {
	data  int
	order int
	jump  *Node
	next  *Node
}

func traverseRec(a *Node, order *int) {
	if a == nil {
		return
	}

	if a.order != -1 {
		return
	}

	a.order = *order
	(*order)++
	fmt.Printf("order %v; data %v\n", a.order, a.data)
	traverseRec(a.jump, order)
	traverseRec(a.next, order)
}

func traverseIterative(a *Node) {
	if a == nil {
		return
	}
	order := 0
	s := stack.NewGenericStack()
	s.Push(interface{}(a))

	fmt.Println("traverseIterative ...")

	for s.Count() > 0 {
		result, _ := s.Pop()
		node := result.(*Node)
		if node.order == -1 {
			node.order = order
			order++
			fmt.Printf("order %v; data %v\n", node.order, node.data)
		} else {
			continue
		}

		if node.next != nil && node.next.order == -1 {
			s.Push(interface{}(node.next))
		}

		if node.jump.order == -1 {
			s.Push(interface{}(node.jump))
		}
	}
}

func main() {
	n1 := &Node{1, -1, nil, nil}
	n2 := &Node{2, -1, nil, nil}
	n3 := &Node{3, -1, nil, nil}
	n4 := &Node{4, -1, nil, nil}
	n1.next = n2
	n2.next = n3
	n3.next = n4
	n1.jump = n3
	n2.jump = n4
	n3.jump = n2
	n4.jump = n4
	//order := 0
	//traverseRec(n1, &order)
	traverseIterative(n1)
}
