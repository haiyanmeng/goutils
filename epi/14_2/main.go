package main

import "fmt"

func mergeSortedArrays(a []int, b []int) []int {
	na, nb := len(a), len(b)
	if nb == 0 {
		return a
	}

	for i := 0; i < nb; i++ {
		a = append(a, 0)
	}

	n := na + nb
	i, j := na-1, nb-1
	for k := n - 1; k >= 0; k-- {
		if i >= 0 && j >= 0 && a[i] >= b[j] {
			a[k] = a[i]
			i--
		} else {
			a[k] = b[j]
			j--
			if j == -1 {
				return a
			}
		}
	}
	return a
}

/*
6 8 9
3 4 5
*/

func main() {
	b1 := []int{5, 6, 7}
	a1 := []int{8, 9, 11, 19}
	fmt.Printf("&a1[0] = %v\n", &a1[0])
	a1 = mergeSortedArrays(a1, b1)
	fmt.Printf("&a1[0] = %v\n", &a1[0])
	fmt.Println(a1)
}
