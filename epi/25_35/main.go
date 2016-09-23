package main

import "fmt"

func f(a []int) int {
	result := 0
	n := len(a)
	i := 0
	for {
		for i < (n-1) && a[i+1] >= a[i] {
			i++
		}
		start := i
		if i >= (n - 1) {
			return result
		}

		i++
		sum := 0
		decreasing := true
		for i <= (n-1) && a[i] < a[start] {
			if a[i] > a[i-1] {
				decreasing = false
			}
			sum += a[i]
			i++
		}
		end := i
		if decreasing {
			continue
		}
		if end > (n - 1) {
			end = n - 1
		}
		delta := end - start - 1
		if delta == 0 {
			continue
		}
		if a[end] < a[start] {
			sum -= a[end]
		}
		result += (delta*min(a[start], a[end]) - sum)
		fmt.Printf("start = %v; end = %v; result = %v\n", start, end, result)
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func main() {
	var a []int
	a = []int{0, 1, 2, 2, 3, 4, 4, 5, 1, 2, 0, 3, 5, 3, 2, 1, 0}
	a = []int{0, 1, 2, 2, 3, 4, 4, 5, 1, 2, 0, 3, 5}
	a = []int{0, 1, 2, 1, 3, 4, 4, 5, 1, 2, 0, 3}
	fmt.Printf("f(%v) = %v\n", a, f(a))
}
