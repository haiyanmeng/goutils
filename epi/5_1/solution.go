package main

import (
	"fmt"
)

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

func main() {
	var x uint64
	x = 0x0b
	fmt.Printf("the parity of %v(%b) is %v\n", x, x, ComputeParity(x))
	
	x = 0x1b
	fmt.Printf("the parity of %v(%b) is %v\n", x, x, ComputeParity(x))

	x = 0x3FFFFFFFFFFFFFFF
	fmt.Printf("the parity of %v(%b) is %v\n", x, x, ComputeParity(x))

	x = 0x7FFFFFFFFFFFFFFF
	fmt.Printf("the parity of %v(%b) is %v\n", x, x, ComputeParity(x))

	x = 0xFFFFFFFFFFFFFFFF
	fmt.Printf("the parity of %v(%b) is %v\n", x, x, ComputeParity(x))
}
