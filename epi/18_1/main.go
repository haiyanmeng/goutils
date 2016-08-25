package main

import (
	"fmt"

	"github.com/hmeng-19/goutils/sort"
)

func alloc(a []int) ([]int, error) {
	n := len(a)
	var result []int
	if n%2 == 1 {
		return result, fmt.Errorf("the job number should be even")
	}

	b := sort.MergeSort(a)
	for i := 0; i < n/2; i++ {
		result = append(result, b[i])
		result = append(result, b[n-i-1])
	}
	return result, nil
}

func main() {
	a, err := alloc([]int{3, 5, 2, 1, 6, 4, 4})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)
}
