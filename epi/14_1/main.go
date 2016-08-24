package main

import "fmt"

func intersect(a []int, b []int) []int {
	p1, p2 := 0, 0
	n1, n2 := len(a), len(b)
	c := []int{}
	n := 0
	for p1 < n1 && p2 < n2 {
		if a[p1] < b[p2] {
			p1++
		} else if a[p1] > b[p2] {
			p2++
		} else {
			if n == 0 || a[p1] != c[n-1] {
				c = append(c, a[p1])
				n++
			}
			p1++
			p2++
		}
	}
	return c
}
func main() {
	a := []int{1, 2, 2, 5, 5, 8, 10}
	b := []int{2, 3, 4, 5, 5, 6, 8, 9, 20}
	fmt.Println(intersect(a, b))

	a = []int{2, 3, 3, 5, 5, 6, 7, 7, 8, 12}
	b = []int{5, 5, 6, 8, 8, 9, 10, 10}
	fmt.Println(intersect(a, b))
}
