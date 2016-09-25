package main

import (
	"fmt"
	"math/rand"
)

func f(n int) int {
	return rand.Intn(n)
}

func randomSubset1(n, k int) []int {
	if k > n || k == 0 {
		return []int{}
	}

	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}

	if k == n {
		return a
	}

	if k > n/2 {
		k = n - k
	}

	for i := 0; i < k; i++ {
		r := f(n - i)
		a[r], a[n-i-1] = a[n-i-1], a[r]
	}
	return a[n-k:]
}

func randomSubset2(n, k int) []int {
	if k > n || k <= 0 {
		return []int{}
	}

	a := []int{}
	if k == n {
		for i := 0; i < k; i++ {
			a = append(a, i)
		}
		return a
	}

	if k > n/2 {
		k = n - k
	}

	m := map[int]int{}
	for i := 0; i < k; i++ {
		r := f(n - i)
		_, ok := m[r]
		if !ok {
			a = append(a, r)
		} else {
			a = append(a, m[r])
		}
		m[r] = n - i - 1
	}
	return a
}

func main() {
	n, k := 100, 10
	for i := 0; i < 20; i++ {
		fmt.Printf("randomSubset1(%v, %v) = %v\n", n, k, randomSubset1(n, k))
		fmt.Printf("randomSubset2(%v, %v) = %v\n", n, k, randomSubset2(n, k))
	}
}
