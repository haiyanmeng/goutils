package main

import "fmt"

func firstIndex(s, t string) int {
	ns, nt := len(s), len(t)
	if ns > nt {
		return -1
	}

	var h int64
	hs := hash(s)
	for i := 0; i < (nt - ns + 1); i++ {
		if i == 0 {
			h = hash(t[:ns])
		} else {
			h = nextHash(h, t[i-1], t[i+ns-1])
		}
		if h == hs && s == t[i:i+ns] {
			return i
		}
	}
	return -1
}

func nextHash(h int64, a, b byte) int64 {
	return h - int64(a) + int64(b)
}

func hash(s string) int64 {
	n := len(s)
	result := int64(0)
	for i := 0; i < n; i++ {
		result += int64(s[i])
	}
	return result
}

func main() {
	s := "hello1"
	t := "good hello hello"
	fmt.Printf("firstIndex(`%v`, `%v`) = %v\n", s, t, firstIndex(s, t))
}
