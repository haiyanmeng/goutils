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

	b := []int{128, 8, 37, 43435, 66, 2343}
	fmt.Println(b)
	sort.BucketSort(b)
	fmt.Println(b)
}
