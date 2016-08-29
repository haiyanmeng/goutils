package main

import "fmt"

func search2D(a [][]int, x int) bool {
	height := len(a)
	if height == 0 {
		return false
	}
	width := len(a[0])
	return searchInPlace(a, 0, height-1, 0, width-1, x)
}

func searchRow(a []int, start, end int, x int) (bool, int) {
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == x {
			return true, mid
		}

		if a[mid] < x {
			if start == end {
				return false, start
			}
			start = mid + 1
		} else {
			if start == end {
				return false, start - 1
			}
			end = mid - 1
		}
	}
	return false, end
}

func searchInPlace(a [][]int, startRow, endRow, startCol, endCol, x int) bool {
	fmt.Printf("searchInPlace(startRow=%v, endRow=%v, startCol=%v, endCol=%v, x=%v)\n", startRow, endRow, startCol, endCol, x)
	if (endRow-startRow) == 0 || (endCol-startCol) == 0 {
		return false
	}
	existed, index := searchRow(a[startRow], startCol, endCol, x)
	fmt.Printf("searchRow(row=%v, startCol=%v, endCol=%v, x=%v) = %v, %v\n", startRow, startCol, endCol, x, existed, index)
	if existed == true {
		return true
	}
	if index == -1 {
		return false
	}
	// remove the startRow
	if index == endCol {
		startRow++
	} else {
		endCol = index
	}
	return searchInPlace(a, startRow, endRow, startCol, endCol, x)
}

func main() {
	a := [][]int{
		[]int{-1, 2, 4, 4, 6},
		[]int{1, 5, 5, 9, 21},
		[]int{3, 6, 6, 9, 22},
		[]int{3, 6, 8, 10, 24},
		[]int{6, 8, 9, 12, 25},
		[]int{8, 10, 12, 13, 40},
	}
	fmt.Println(search2D(a, 7))
	fmt.Println(search2D(a, 8))
	fmt.Println(search2D(a, -100))
	fmt.Println(search2D(a, 100))
}
