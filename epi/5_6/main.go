package main

import "fmt"

func div(a, b int) int {
	r := 0
	for a >= b {
		d := b
		e := 1
		for (d << 1) <= a {
			e <<= 1
			d <<= 1
		}
		r += e
		a -= d
	}
	return r
}

func div1(a, b int) int {
	r := 0
	if b > a {
		return r
	}

	e := 1
	d := b
	for (d << 1) <= a {
		e <<= 1
		d <<= 1
	}
	r += e
	a -= d

	e >>= 1
	d >>= 1
	for e > 0 {
		if d <= a {
			r += e
			a -= d
		}
		e >>= 1
		d >>= 1
	}
	return r
}
func main() {
	fmt.Println(div(35454, 7))
	fmt.Println(div1(35454, 7))
}
