package main

import "fmt"

func max(x1, x2, x3, x4 int) int {
	m := x1
	if x2 > m {
		m = x2
	}
	if x3 > m {
		m = x3
	}
	if x4 > m {
		m = x4
	}
	return m
}

func f(a []int, start, end int) int {
	n := end - start + 1
	if n == 0 {
		return 0
	}

	x1 := f(a, 2, n-1) + a[start]
	x2 := f(a, 1, n-2) + a[start]
	x3 := f(a, 1, n-2) + a[end]
	x4 := f(a, 0, n-3) + a[end]
	return max(x1, x2, x3, x4)
}

func fCache(a []int, start, end int, cache map[string]int) int {
	n := end - start + 1
	if n == 0 {
		return 0
	}

	key := fmt.Sprintf("%d %d", start, end)
	if value, ok := cache[key]; ok {
		return value
	}

	x1 := fCache(a, 2, n-1, cache) + a[start]
	x2 := fCache(a, 1, n-2, cache) + a[start]
	x3 := fCache(a, 1, n-2, cache) + a[end]
	x4 := fCache(a, 0, n-3, cache) + a[end]
	m := max(x1, x2, x3, x4)
	cache[key] = m
	return m
}

func main() {
	a := []int{25, 5, 10, 5, 10, 5, 10, 25, 1, 25, 1, 25, 1, 25, 5, 10}
	cache := map[string]int{}
	fmt.Printf("f: %v\n", f(a, 0, len(a)-1))
	fmt.Printf("fCache: %v\n", fCache(a, 0, len(a)-1, cache))
	fmt.Printf("len(cache) = %v; %#v\n", len(cache), cache)
}
