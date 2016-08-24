package main

import "fmt"

func isP(a string) bool {
	count := map[byte]int{}
	n := len(a)
	for i := 0; i < n; i++ {
		count[a[i]]++
	}

	odds := 0
	for _, v := range count {
		if v%2 == 1 {
			odds++
		}
	}
	return odds <= 1
}

func main() {
	s := "level"
	fmt.Printf("isP(%v) = %v\n", s, isP(s))
	s = "dlevel"
	fmt.Printf("isP(%v) = %v\n", s, isP(s))
}
