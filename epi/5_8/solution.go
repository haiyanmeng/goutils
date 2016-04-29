package main

import (
	"fmt"
)

func reverseDigits(x int) int {
	neg := false
	if x < 0 {
		neg = true
		x *= -1
	}

	i := 0
	for x != 0 {
		i = i * 10 + x%10
		x /= 10
	}

	if neg {
		i *= -1
	}
	return i
}

func reverseDigitsTest(x int) {
	fmt.Printf("reverseDigits(%d) = %d\n", x, reverseDigits(x))
}

func main() {
	reverseDigitsTest(100);	
	reverseDigitsTest(-34347551);	
}
