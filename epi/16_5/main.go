package main

import (
	"fmt"
)

func calCombinations(a []int, start, end int, k int) [][]int {
	n := end - start + 1

	if k == 0 || n < k {
		return [][]int{}
	}

	if n == k {
		return [][]int{a[start : start+n]}
	}

	r := [][]int{}
	for i := 0; i <= n-k; i++ {
		r2 := calCombinations(a, start+i+1, end, k-1)

		if len(r2) == 0 {
			r1 := []int{a[start+i]}
			r = append(r, r1)
			continue
		}

		for _, item := range r2 {
			r1 := []int{a[start+i]}
			r1 = append(r1, item...)
			r = append(r, r1)
		}
	}
	return r
}

func f(a []int, start, end int, k int) [][]int {
	n := end - start + 1

	if k == 0 || n < k {
		return [][]int{}
	}

	if n == k {
		return [][]int{a[start : start+n]}
	}

	r := [][]int{}

	r1 := f(a, start+1, end, k)
	r2 := f(a, start+1, end, k-1)

	r = append(r, r1...)

	if len(r2) == 0 {
		r = append(r, []int{a[start]})
		return r
	}

	for _, item := range r2 {
		r3 := []int{a[start]}
		r3 = append(r3, item...)
		r = append(r, r3)
	}
	return r

}

func main() {
	a := []int{11, 12, 13, 14, 15}
	k := 2
	r := calCombinations(a, 0, len(a)-1, k)
	fmt.Println(r)
	fmt.Println(f(a, 0, len(a)-1, k))
}
