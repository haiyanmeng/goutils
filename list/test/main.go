package main

import (
	"github.com/hmeng-19/goutils/list"
)

func main() {
	list := list.New([]int{1, 2, 3, 4, 5, 6})
	list.Traverse()
}
