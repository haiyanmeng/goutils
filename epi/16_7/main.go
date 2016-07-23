package main

import (
	"fmt"
)

func isP(s string) bool {
	if s == "" {
		return true
	}

	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

func calPD(s string) [][]string {
	if s == "" {
		return [][]string{}
	}

	n := len(s)
	if n == 1 {
		return [][]string{[]string{s}}
	}

	r := [][]string{}
	for i := 0; i < n; i++ {
		if !isP(s[:i+1]) {
			continue
		}
		r1 := calPD(s[i+1:])
		if len(r1) == 0 {
			r = append(r, []string{s[:i+1]})
		}

		for _, item := range r1 {
			r2 := []string{s[:i+1]}
			r2 = append(r2, item...)
			r = append(r, r2)
		}
	}
	return r
}
func main() {
	s := "611116"
	r := calPD(s)
	fmt.Println(r)
}
