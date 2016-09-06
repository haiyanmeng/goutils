package main

import "fmt"

/*
4 5 0 3 2 1
4 5 1 0 2 3

4 0 3 2 5 1
4 0 3 5


4 3 2 1 0
from backward, find a[i-1] < a[i]
*/
func nextP(b []int) []int {
	n := len(b)
	if n < 2 {
		return b
	}

	a := make([]int, 0, n)
	for j := 0; j < n; j++ {
		a = append(a, b[j])
	}

	i := findDecrease(a, n)
	if i == -1 {
		return a
	}

	j := findLargerFromEnd(a, a[i], n)
	a[i], a[j] = a[j], a[i]
	swapRange(a, i+1, n-1)
	return a
}

func swapRange(a []int, start, end int) {
	i, j := start, end
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

func findLargerFromEnd(a []int, target int, n int) int {
	for i := n - 1; i >= 0; i-- {
		if a[i] > target {
			return i
		}
	}
	return -1
}

func findDecrease(a []int, n int) int {
	for i := n - 1; i >= 0; i-- {
		if i == 0 {
			return -1
		}

		if a[i] > a[i-1] {
			return i - 1
		}
	}
	return -1
}

func main() {
	a := []int{1, 0, 3, 2}
	a = []int{4, 5, 0, 3, 2, 1}
	b := nextP(a)
	fmt.Printf("nextP(%v) = %v\n", a, b)
}
