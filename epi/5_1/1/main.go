package main

import (
	"fmt"
)

func parity1(a int64) string {
	odd := 0
	for i := 0; i < 64; i++ {
		if a&(0x01<<uint(i)) != 0 {
			odd ^= 1
		}
	}

	if odd == 1 {
		return "odd"
	}
	return "even"
}

func parity(a int64) string {
	odd := 0
	for a != 0 {
		odd ^= 1
		a &= (a - 1)
	}

	if odd == 1 {
		return "odd"
	}
	return "even"
}

func main() {
	a := int64(80)
	a = 1024
	fmt.Printf("parity(%v) = %v\n", a, parity(a))
	a = 10345
	a = 9223372036854775807
	fmt.Printf("parity(%v) = %v\n", a, parity(a))
}
