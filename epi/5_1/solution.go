package main

import (
	"fmt"
)

// time complexity O(n), n is the word width in bits
func ComputeParity(x uint64) (r bool) { // the (r bool) must be surrounded by a pair of parenthess
	var i uint
	for i=0; i<64; i++ {
		if x & (0x01<<i) != 0 {
			if r {
				r = false
			} else {
				r = true
			}
		}
	}
	return r
}

// time complexity is O(d), d is the number of the bits which are set to be 1
func ComputeParity1(x uint64) (r bool) {
	var c uint = 0
	for x != 0 {
		x &= (x-1)
		c += 1
	}

	if c%2 == 0 {
		return false
	} else {
		return true
	}
}

type FuncTemp func(x uint64) (r bool)

func test(f FuncTemp) {
	var x uint64
	x = 0x0b
	fmt.Printf("the parity of %v(%b) is %v\n", x, x, f(x))
	
	x = 0x1b
	fmt.Printf("the parity of %v(%b) is %v\n", x, x, f(x))

	x = 0x3FFFFFFFFFFFFFFF
	fmt.Printf("the parity of %v(%b) is %v\n", x, x, f(x))

	x = 0x7FFFFFFFFFFFFFFF
	fmt.Printf("the parity of %v(%b) is %v\n", x, x, f(x))

	x = 0xFFFFFFFFFFFFFFFF
	fmt.Printf("the parity of %v(%b) is %v\n", x, x, f(x))
}

func main() {
	test(ComputeParity)
	test(ComputeParity1)
}
