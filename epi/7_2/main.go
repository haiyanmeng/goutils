package main

import (
	"fmt"
	"os"
)

/*
123 base 10
b2 = 8
123/8 = 15
123%8 = 3
15/8 = 1
15%8 = 7
173
64+56+3=
*/

func digit2Char(x int) string {
	var b byte
	if x < 10 {
		b = '0' + byte(x)
	} else {
		b = 'A' + byte(x-10)
	}
	return string([]byte{b})
}

func int2Str(x int, base byte) string {
	var s string
	for x > 0 {
		reminder := x % int(base)
		s = digit2Char(reminder) + s
		x = x / int(base)
	}
	return s
}

func validateChar(b byte, base byte) (bool, byte) {
	diff := b - '0'
	if diff < 0 {
		return false, 0
	}

	if base <= 10 {
		if diff >= base {
			return false, 0
		}
		return true, diff
	} else {
		if diff < 10 {
			return true, diff
		}

		diffA := b - 'A'
		if diffA >= 0 && diffA <= 5 {
			return true, 10 + diffA
		}
		return false, 0
	}
}

func str2Int(s string, base byte) (int, error) {
	r := 0
	n := len(s)
	factor := 1
	for i := 0; i < n; i++ {
		k, v := validateChar(s[n-i-1], base)
		if !k {
			return 0, fmt.Errorf("%c is not valid character for base %d\n", s[n-i-1], base)
		}
		r += int(v) * factor
		factor *= int(base)
	}
	return r, nil
}

func checkBase(base byte) bool {
	if base >= 2 && base <= 16 {
		return true
	}
	return false
}

func baseConversion(s string, b1, b2 byte) (string, error) {
	if !checkBase(b1) || !checkBase(b2) {
		return "", fmt.Errorf("the base should be within [2, 16]")
	}

	x, err := str2Int(s, b1)
	fmt.Printf("str2Int(`%s`, %v)=%v\n", s, b1, x)
	if err != nil {
		return "", err
	}

	return int2Str(x, b2), nil
}

func main() {
	s1 := "615"
	var b1 byte = 7
	var b2 byte = 13
	s2, err := baseConversion(s1, b1, b2)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("baseConversion(`%s`, %d, %d) = `%s`\n", s1, b1, b2, s2)
}
