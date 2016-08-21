package main

import (
	"fmt"
)

func increment(a *[]int) {
	if (*a)[0] < 9 {
		(*a)[0]++
		return
	}

	extra := 0
	b := 0
	for i, k := range *a {
		if i == 0 {
			b = 1
		} else {
			b = 0
		}
		(*a)[i] = (k + b + extra) % 10
		extra = (k + b + extra) / 10
		if extra == 0 {
			return
		}
	}

	if extra == 1 {
		*a = append(*a, 1)
	}
}

func main() {
	var a []int
	a = []int{3, 4, 5} //543
	fmt.Println(a)
	increment(&a)
	fmt.Println(a)

	a = []int{9, 9, 9, 8} //543
	fmt.Println(a)
	increment(&a)
	fmt.Println(a)

	a = []int{9, 9, 9} //543
	fmt.Println(a)
	increment(&a)
	fmt.Println(a)
}
