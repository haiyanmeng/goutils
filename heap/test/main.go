package main

import (
	"fmt"

	"github.com/hmeng-19/goutils/heap"
)

func main() {
	h := heap.New([]int{6, 2, 8, 9, 10, 3, 1})
	h.Print()

	for {
		top, err := h.Delete()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("the top is %v, the count is: %v\n", top, h.Count())
		h.Print()
	}
}
