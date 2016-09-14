package main

import "fmt"

func findMin(a []int) int {
	n := len(a)
	i := 0
	for i < n {
		if a[i] < 1 || a[i] > n {
			i++
			continue
		}
		x := a[i]
		if i == x-1 ||
			a[i] == a[x-1] {
			i++
			continue
		}
		a[i], a[x-1] = a[x-1], a[i]
	}

	for i := 0; i < n; i++ {
		if a[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	var a []int
	a = []int{-1, -3, 30, 3, 5, 4, 5, 10}
	fmt.Println(a)
	fmt.Printf("findMin(%v) = %v\n", a, findMin(a))
}
