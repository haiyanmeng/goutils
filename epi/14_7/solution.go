package main

import (
	"fmt"
	"sort"
)

type endpoint struct {
	d    int
	open bool
}

type interval struct {
	left  endpoint
	right endpoint
}

type intervals []interval

func (in intervals) Len() int {
	return len(in)
}

func (in intervals) Less(i, j int) bool {
	return in[i].left.d <= in[j].left.d
}

func (in intervals) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

func openUnion(s, t bool) bool {
	if s == true && t == true {
		return true
	}
	return false
}

func merge(a, b interval) interval {
	var left, right endpoint

	if a.left.d < b.left.d {
		left = a.left
	} else {
		// a.left.d == b.left.d
		left.d = a.left.d
		left.open = openUnion(a.left.open, b.left.open)
	}

	if a.right.d < b.right.d {
		right = b.right
	} else if a.right.d > b.right.d {
		right = a.right
	} else {
		right.d = a.right.d
		right.open = openUnion(a.right.open, b.right.open)
	}

	return interval{left, right}
}

func mergeIntervals(in intervals) intervals {
	out := []interval{}
	n := len(in)
	if n == 0 {
		return intervals(out)
	}

	cur := in[0]
	for i := 1; i < n; i++ {
		next := in[i]
		if cur.right.d < next.left.d ||
			(cur.right.d == next.left.d && cur.right.open == next.left.open == true) {
			out = append(out, cur)
			cur = next
		} else {
			// merge cur and next
			cur = merge(cur, next)
		}
	}
	out = append(out, cur)
	return intervals(out)
}

func main() {
	s := intervals([]interval{
		{endpoint{0, true}, endpoint{3, true}},
		{endpoint{1, false}, endpoint{1, false}},
		{endpoint{2, false}, endpoint{4, false}},
		{endpoint{3, false}, endpoint{4, true}},
		{endpoint{12, false}, endpoint{14, false}},
		{endpoint{8, false}, endpoint{11, true}},
		{endpoint{12, true}, endpoint{16, false}},
		{endpoint{9, false}, endpoint{11, false}},
		{endpoint{16, true}, endpoint{17, true}},
		{endpoint{5, false}, endpoint{7, true}},
		{endpoint{7, false}, endpoint{8, true}},
		{endpoint{13, true}, endpoint{15, true}},
	})

	sort.Sort(s)

	fmt.Println(s)

	fmt.Println("\n\n")

	fmt.Println(mergeIntervals(s))
}
