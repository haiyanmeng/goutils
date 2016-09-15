package main

import "fmt"

func rotate(a []int, n int, i int) {
	if n <= 1 {
		return
	}
	i = i % n
	if i == 0 {
		return
	}

	j := n - i
	if j >= i {
		swap(a, j-i, j, i)
		rotate(a, j, i)
	} else {
		swap(a, 0, n-j, j)
		rotate(a, i, i-j)
	}
}

func swap(a []int, s, t, len int) {
	for i := 0; i < len; i++ {
		a[s+i], a[t+i] = a[t+i], a[s+i]
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 4
	fmt.Printf("before rotating: %v\n", a)
	rotate(a, len(a), i)
	fmt.Printf("after rotating: %v\n", a)
}
