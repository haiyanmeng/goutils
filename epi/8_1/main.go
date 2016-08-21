package main

import (
	"fmt"

	"github.com/hmeng-19/goutils/list"
)

func mergeLists(list1, list2 *list.List) *list.List {
	var list4 *list.List // this is not okay: list3 is (*list.List)(nil)
	fmt.Printf("list4 = %v; list4 == nil? %v\n", list4, list4 == nil)
	list3 := &list.List{} // this works. list3 is &list.List{Head:(*list.Node)(nil), Tail:(*list.Node)(nil)}
	//list3 := list.New([]int{}) // this also works
	fmt.Printf("%#v\n", list3)
	fmt.Printf("list3 == nil? %v\n", list3 == nil)
	p1 := list1.Head
	p2 := list2.Head
	for p1 != nil || p2 != nil {
		if p1 == nil {
			list3.AddNode(p2)
			list3.Tail = list2.Tail
			break
		}

		if p2 == nil {
			list3.AddNode(p1)
			list3.Tail = list1.Tail
			break
		}

		if p1.Data <= p2.Data {
			list3.AddNode(p1)
			p1 = p1.Next
		} else {
			list3.AddNode(p2)
			p2 = p2.Next
		}
	}

	return list3
}

func main() {
	list1 := list.New([]int{1, 3, 5, 7, 9})
	list2 := list.New([]int{2, 4, 6, 8, 10})
	list3 := mergeLists(list1, list2)
	list3.Traverse()
}
