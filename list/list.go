package list

import "fmt"

type List struct {
	Head, Tail *Node
}

type Node struct {
	Data int
	Next *Node
}

func New(a []int) *List {
	list := &List{}
	for _, elem := range a {
		list.Add(elem)
	}
	return list
}

func (list *List) Add(a int) {
	node := &Node{a, nil}
	list.AddNode(node)
}

func (list *List) AddNode(node *Node) {
	if list.Head == nil {
		list.Head = node
	} else {
		list.Tail.Next = node
	}
	list.Tail = node
}

func (list *List) Traverse() {
	p := list.Head
	for p != nil {
		fmt.Println(p.Data)
		p = p.Next
	}
}
