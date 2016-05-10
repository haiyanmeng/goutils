package main

import (
	"fmt"
	"unsafe"
)

func getEnd(a []int) bool {
	n := int(unsafe.Sizeof(a))
	d := 0 
	for i, x := range a {
		if i > d {
			return false
		}

		if (i+x) >= (n-1) {
			return true
		}

		if (i+x) > d {
			d = i+x
		}
	}
	return true 
}

func getEndTest(a []int) {
	fmt.Printf("getEnd(%v) = %v\n", a, getEnd(a))
}

func main() {
	a := []int{3, 3, 1, 0, 2, 0, 1}
	getEndTest(a)
	a = []int{3, 2, 1, 0, 2, 0, 1}
	getEndTest(a)
}
