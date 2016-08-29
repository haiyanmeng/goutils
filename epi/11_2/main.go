package main

import "fmt"

func reverse(a []int, start, end int) {
	for start < end {
		a[start], a[end] = a[end], a[start]
		start++
		end--
	}
}

func convertDecrease(a []int, k int) {
	n := len(a)
	cur := 0
	for i := 0; i < k; i++ {
		if cur >= (n - 1) {
			return
		}
		start := cur

		inc := false
		if a[cur] < a[cur+1] {
			inc = true
			for cur < (n-1) && a[cur] < a[cur+1] {
				cur++
			}
		} else {
			for cur < (n-1) && a[cur] > a[cur+1] {
				cur++
			}
		}
		end := cur
		cur++
		fmt.Printf("inc: %v; start: %v; end: %v\n", inc, start, end)
		if inc {
			continue
		}
		reverse(a, start, end)
	}
}

func main() {
	a := []int{57, 131, 493, 294, 339, 418, 452, 442, 190}
	fmt.Println(a)
	convertDecrease(a, 4)
	fmt.Println(a)
}
