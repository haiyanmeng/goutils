package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func count1(n1, n2 int) int {
	if n1 <= 0 || n2 <= 0 {
		return 0
	}

	if n1 == 1 || n2 == 1 {
		return 1
	}

	r := count1(n1, n2-1) + count1(n1-1, n2)
	return r
}

func count(n1, n2 int, cache [][]int) int {
	if n1 <= 0 || n2 <= 0 {
		return 0
	}

	if n1 == 1 || n2 == 1 {
		cache[n1-1][n2-1] = 1
		return 1
	}

	if cache[n1-1][n2-1] != 0 {
		return cache[n1-1][n2-1]
	}

	r := count(n1, n2-1, cache) + count(n1-1, n2, cache)
	cache[n1-1][n2-1] = r
	return r
}

func init2DArray(n1, n2 int) [][]int {
	r := make([][]int, n1)
	for i:=0; i<n1; i++ {
		r[i] = make([]int, n2)
	}
	return r
}

func main() {
	n1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	n2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	t1 := time.Now()
	cache := init2DArray(n1, n2)
	fmt.Printf("count(%d, %d) = %d\n", n1, n2, count(n1, n2, cache))
	t2 := time.Now()
	fmt.Printf("count takes %v\n", t2.Sub(t1))

	fmt.Printf("count1(%d, %d) = %d\n", n1, n2, count1(n1, n2))
	fmt.Printf("count1 takes %v\n", time.Now().Sub(t2))
}
