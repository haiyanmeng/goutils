package main

import "fmt"

func add(a, b int) int {
	//fmt.Printf("add(a, b) a = %32b; b = %32b\n", a, b)
	c := a & b
	d := a | b
	//fmt.Printf("c = %32b; d = %32b\n", c, d)
	k := 0x01

	for {
		if c == 0 {
			return d
		}

		if c&0x01 == 1 {

			e := k
			for {
				if d&e == 0 {
					d ^= e
					break
				}
				d ^= e
				e <<= 1
			}
		}

		c >>= 1
		k <<= 1
		//fmt.Printf("c = %32b; k = %32b; d = %32b\n", c, k, d)
	}

	return d
}

func mul(a, b int) int {
	result := 0
	for {
		if b == 0 {
			return result
		}

		if b&0x01 != 0 {
			result = add(result, a)
		}
		a <<= 1
		b >>= 1
	}
}

func main() {
	var a, b int
	a, b = 15, 25
	fmt.Printf("mul(%v, %v) = %v; %v * %v = %v\n", a, b, mul(a, b), a, b, a*b)

	a, b = 158, 275
	fmt.Printf("mul(%v, %v) = %v; %v * %v = %v\n", a, b, mul(a, b), a, b, a*b)

	a, b = 13, 2583
	fmt.Printf("mul(%v, %v) = %v; %v * %v = %v\n", a, b, mul(a, b), a, b, a*b)
}
