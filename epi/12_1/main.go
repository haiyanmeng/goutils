package main

import "fmt"

func binarySearch(a []int, target int) int {
	left, right := 0, len(a)-1
	for left <= right {
		mid := (left + right) / 2
		if a[mid] == target {
			return mid
		}

		if a[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func binarySearchMin(a []int, target int) int {
	left, right := 0, len(a)-1
	last := -1
	for left <= right {
		mid := (left + right) / 2
		if a[mid] == target {
			last = mid
			right = mid - 1
		} else if a[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return last
}

func main() {
	a := []int{1, 2, 3, 3, 4, 5, 6, 8, 8, 8, 9, 10}
	target := 7
	fmt.Printf("binarySearch(%v, %v) = %v\n", a, target, binarySearch(a, target))
	target = 8
	fmt.Printf("binarySearch(%v, %v) = %v\n", a, target, binarySearch(a, target))
	target = 7
	fmt.Printf("binarySearchMin(%v, %v) = %v\n", a, target, binarySearchMin(a, target))
	target = 8
	fmt.Printf("binarySearchMin(%v, %v) = %v\n", a, target, binarySearchMin(a, target))
}
