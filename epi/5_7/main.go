package main

import "fmt"

func pow(x float64, y int) float64 {
	if x == 0 {
		return 0
	}
	if y < 0 {
		x = 1 / x
		y = -y
	}
	var result float64 = 1
	factor := x
	for y > 0 {
		if y&0x01 == 0x01 {
			result *= factor
		}
		factor *= factor
		y = y >> 1
	}
	return result
}

func main() {
	x := 2.0
	y := -7
	fmt.Printf("pow(%v, %v) = %v\n", x, y, pow(x, y))
}
