package main

import (
	"fmt"
)

func two_sum(s []int, target int) (i1, i2 int) {
	n := len(s)
	for i:=0; i<(n-1); i++ {
		m := target - s[i]
		for j:=i+1; j<n; j++ {
			if s[j] == m {
				i1, i2 = i, j
				return
			}
		}
	}
	return
}

func two_sum_1(s []int, target int) (i1, i2 int) {
	d := make(map[int]int) // key is the element of s; value is the index of each element 
	for i, k := range s {
		d[k] = i
	}

	n := len(s)
	for i:=0; i<(n-1); i++ {
		if v, ok := d[target - s[i]]; ok {
			i1 = i
			i2 = v
			break
		} 
	}
	return
}

func main() {
	var s []int
	var target int
	s = []int{2, 11, 8, 7, 15}
	target = 9
	i1, i2 := two_sum(s, target)
	fmt.Printf("s[%d](%d) + s[%d](%d) = %d\n", i1, s[i1], i2, s[i2], target)
	i1, i2 = two_sum_1(s, target)
	fmt.Printf("s[%d](%d) + s[%d](%d) = %d\n", i1, s[i1], i2, s[i2], target)
}
