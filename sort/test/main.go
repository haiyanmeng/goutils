package main

import (
	"fmt"

	"github.com/hmeng-19/goutils/sort"
)

func main() {
	fmt.Println(sort.MergeSort([]int{5, 2, 1, 6, 4, 4}))
	a := []int{5, 2, 1, 6, 4, 4}
	a = []int{5, 10, 5, 6, 3, 2, 1}
	a = []int{5, 6, 6}
	sort.QuickSort(a)
	fmt.Println(a)
}
