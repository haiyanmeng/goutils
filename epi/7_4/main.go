package main

import "fmt"

func f(s []byte, size int) {
	t := removeB(s, size)
	n := countA(s, t)
	replaceA(s, t, t+n)
}

func replaceA(s []byte, size int, t int) {
	j := t - 1
	for i := size - 1; i >= 0; i-- {
		if s[i] == 'a' {
			s[j] = 'd'
			s[j-1] = 'd'
			j -= 2
		} else {
			s[j] = s[i]
			j--
		}
	}
}

func countA(s []byte, size int) int {
	n := 0
	for i := 0; i < size; i++ {
		if s[i] == 'a' {
			n++
		}
	}
	return n
}

func removeB(s []byte, size int) int {
	j := 0
	for i := 0; i < size; i++ {
		if s[i] != 'b' {
			s[j] = s[i]
			j++
		}
	}
	return j
}

func main() {
	//t := []byte{'a', 'c', 'd', 'b', 'b', 'c', 'a'}
	t := []byte("abac")
	fmt.Println(t)
	s := make([]byte, 10000)
	n := len(t)
	for i := 0; i < n; i++ {
		s[i] = t[i]
	}
	f(s, n)
	fmt.Println(string(s))
}
