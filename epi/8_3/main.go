package main

import (
	"fmt"

	"github.com/hmeng-19/goutils/list"
)

func hasCycle(list1 *list.List) *list.Node {
	visited := map[*list.Node]bool{}
	p := list1.Head
	for p != nil {
		if _, ok := visited[p]; ok {
			return p
		}
		visited[p] = true
		p = p.Next
	}
	return nil
}

func hasCycle1(list1 *list.List) *list.Node {
	// test whether there is a cycle
	p1, p2 := list1.Head, list1.Head
	for {
		if p1 == nil || p2 == nil {
			return nil
		}
		p1 = p1.Next
		p2 = p2.Next
		if p2 == nil {
			return nil
		}
		p2 = p2.Next
		if p1 == p2 {
			break
		}
	}

	// get the size of the cycle
	size := 1
	for {
		p2 = p2.Next
		if p2 == p1 {
			break
		}
		size++
	}

	// find the start of the cycle
	p1, p2 = list1.Head, list1.Head
	for i := 0; i < size; i++ {
		p2 = p2.Next
	}
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}

func main() {
	list1 := list.New([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(hasCycle(list1))
	fmt.Println(hasCycle1(list1))

	list2 := list.New([]int{1, 2, 3, 4, 5, 6})
	list2.AddNode(list2.Head.Next.Next)
	fmt.Println(hasCycle(list2))
	fmt.Println(hasCycle1(list2))
}
