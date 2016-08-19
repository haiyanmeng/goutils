package main

import "fmt"

func int2Str(a int) string {
	t := []byte{}
	neg := false
	if a < 0 {
		a = -a
		neg = true
	}

	for a > 0 {
		t = append(t, byte('0'+a%10))
		a /= 10
	}

	result := []byte{}
	if neg {
		result = append(result, byte('-'))
	}

	n := len(t)
	for i := n - 1; i >= 0; i-- {
		result = append(result, t[i])
	}
	return string(result)
}

func str2Int(a string) (int, error) {
	n := len(a)
	if n == 0 {
		return 0, fmt.Errorf("the string is empty")
	}

	i := 0
	for a[i] == ' ' {
		i++
	}

	if a[i] == '+' {
		i++
	}

	neg := false
	if a[i] == '-' {
		neg = true
		i++
	}

	result := 0
	for i < n {
		if a[i] < '0' || a[i] > '9' {
			return 0, fmt.Errorf("the string is not a legal integer")
		}
		c := a[i] - '0'
		result = result*10 + int(c)
		i++
	}

	if neg {
		result *= -1
	}

	return result, nil
}

func main() {
	a := 566
	a = -34354543
	fmt.Printf("int2Str(%d) = %s\n", a, int2Str(a))
	b := "+566"
	b = "   - 63435656"
	d, err := str2Int(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("str2Int(%s) = %d\n", b, d)
}
