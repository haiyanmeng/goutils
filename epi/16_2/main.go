package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func hasConflict(a []int, row int) bool {
	for i := 0; i < row; i++ {
		if a[i] == a[row] || // there are two queens on a single column
			abs(row-i) == abs(a[row]-a[i]) { // there are two queens on a diagonal
			return true
		}
	}
	return false
}

func AllPlacements(a []int, n int, row int, result *[][]int) {
	fmt.Printf("a=%v row=%v\n", a, row)
	if row < 0 {
		return
	}

	// find a valid placement
	if row == n {
		b := []int{}
		b = append(b, a...)
		*result = append(*result, b)
		fmt.Printf("a=%v row=%v  result=%v\n", a, row, *result)
		row--
	}

	a[row]++

	if a[row] == n {
		if row == 0 {
			return
		}
		a[row] = -1 // reset a[row] to be -1;
		AllPlacements(a, n, row-1, result)
	} else {
		if hasConflict(a, row) {
			AllPlacements(a, n, row, result)
		} else {
			AllPlacements(a, n, row+1, result)
		}
	}
}

func main() {
	n := 8
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		a[i] = -1
	}

	r := [][]int{}
	AllPlacements(a, n, 0, &r)
	for i, placement := range r {
		fmt.Printf("The %dth placement:\n", i+1)
		fmt.Println(placement)
	}
}
