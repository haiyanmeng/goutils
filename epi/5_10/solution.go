package main

import (
	"fmt"
	"math/rand"
)

func randGen(a, b int) int {
	n := b - a + 1 // the total number of integers of between [a, b] inclusively

	m := 1
	k := 0
	for k=0; m<n; k++ {
		m *= 2
	}

	// m >= n, k bits are needed to express n integers
	r := n
	var t int64 = 1
	for r < 0 || r >= n {
		fmt.Printf("running once ... m = %d; r = %d\n", m, r)
		r = 0
		rand.Seed(t)
		t++
		for i:=0; i<k; i++ {
			r |= rand.Intn(2) << uint(i)
		}
	}

	return r+a
}

func randGenTest(a, b int) {
	fmt.Printf("randGen(%d, %d) = %d\n", a, b, randGen(a, b))
}

func main() {
	randGenTest(10, 20)
	randGenTest(10, 200000)
}
