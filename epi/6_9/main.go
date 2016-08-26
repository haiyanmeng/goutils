package main

import (
	"fmt"
	"log"
)

func permute(a []string, b []int) []string {
	n := len(a)
	if len(b) != n {
		log.Fatal("the length of the array and the permutation should be the same")
	}
	c := make([]string, n)
	for i := 0; i < n; i++ {
		c[b[i]] = a[i]
	}
	return c
}

func permuteInPlace(a []string, b []int) {
	n := len(a)
	if len(b) != n {
		log.Fatal("the length of the array and the permutation should be the same")
	}

	for i := 0; i < n; i++ {
		b[i] -= i
	}

	for i := 0; i < n; i++ {
		for b[i] != 0 {
			j := i + b[i]
			a[i], a[j] = a[j], a[i]
			b[i] = b[j] + b[i]
			b[j] = 0
		}
	}
}
func main() {
	a := []string{"a", "b", "c", "d"}
	b := []int{2, 0, 1, 3}
	fmt.Printf("before permute, a: %v; b: %v\n", a, b)
	c := permute(a, b)
	fmt.Printf(" after permute, a: %v; b: %v; c: %v\n", a, b, c)

	fmt.Printf("before permuteInPlace, a: %v; b: %v\n", a, b)
	permuteInPlace(a, b)
	fmt.Printf(" after permuteInPlace, a: %v; b: %v\n", a, b)
}
