package main

import (
	"fmt"
	"time"
)

func dist1(s1, s2 string) int {
	n1, n2 := len(s1), len(s2)
	
	if n1 == 0 && n2 == 0 {
		return 0
	}

	if n1 == 0 && n2 != 0 {
		return n2
	}

	if n1 != 0 && n2 == 0 {
		return n1
	}

	if s1[n1-1] == s2[n2-1] {
		r := dist1(s1[:n1-1], s2[:n2-1])
		return r
	}

	r1 := dist1(s1[:n1-1], s2[:n2-1])

	r2 := dist1(s1[:n1-1], s2)

	r3 := dist1(s1, s2[:n2-1])

	min := r1
	if r2 < min {
		min = r2
	}

	if r3 < min {
		min = r3
	}

	return min + 1
}

func dist(s1, s2 string, d map[string]int) int {
	n1, n2 := len(s1), len(s2)
	key := fmt.Sprintf("%d,%d", n1, n2)
	if v, ok := d[key]; ok {
		return v
	}

	if n1 == 0 && n2 == 0 {
		d[key] = 0
		return 0
	}

	if n1 == 0 && n2 != 0 {
		d[key] = n2
		return n2
	}

	if n1 != 0 && n2 == 0 {
		d[key] = n1
		return n1
	}

	if s1[n1-1] == s2[n2-1] {
		r := dist(s1[:n1-1], s2[:n2-1], d)
		d[key] = r
		return r
	}

	r1 := dist(s1[:n1-1], s2[:n2-1], d)

	r2 := dist(s1[:n1-1], s2, d)

	r3 := dist(s1, s2[:n2-1], d)

	min := r1
	if r2 < min {
		min = r2
	}

	if r3 < min {
		min = r3
	}

	d[key] = min+1
	return min + 1
}

func main() {
	s1 := "Saturdaydf343rfdfdg4rgfjjjjjjjjjjjjjjjjjjferererewrew"
	s2 := "Sundaysdfdfsafdfsaaaaaaaaaaaaaaaadfqrewriwetieiwtiwtiewtiwirwir234fd"
	//s1 = "Saturday"
	//s2 = "Sundays"
	d := make(map[string]int)
	t1 := time.Now()
	fmt.Printf("dist(%s, %s) = %d\n", s1, s2, dist(s1, s2, d))
	t2 := time.Now()
	fmt.Printf("dist with cache takes %v\n", t2.Sub(t1))
	fmt.Printf("dist1(%s, %s) = %d\n", s1, s2, dist1(s1, s2))
	t3 := time.Now()
	fmt.Printf("dist1 takes %v\n", t3.Sub(t2))
}
