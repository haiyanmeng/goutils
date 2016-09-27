package main

import "fmt"

type List struct {
	data int
	next *List
}

func f(a, b *List) *List {
	if a == nil || b == nil {
		return nil
	}
	na, lasta := lastNode(a)
	nb, lastb := lastNode(b)
	if lasta != lastb {
		return nil
	}

	pa, pb := a, b
	if na >= nb {
		for i := 0; i < (na - nb); i++ {
			pa = pa.next
		}
	} else {
		for i := 0; i < (nb - na); i++ {
			pb = pb.next
		}
	}
	for pa != pb {
		pa, pb = pa.next, pb.next
	}
	return pa
}

func lastNode(a *List) (int, *List) {
	n := 0
	p := a
	var last *List = nil
	for p != nil {
		n++
		last = p
		p = p.next
	}
	return n, last
}

func main() {
	n1 := &List{1, nil}
	n2 := &List{2, nil}
	n3 := &List{3, nil}
	n4 := &List{4, nil}
	n5 := &List{5, nil}
	n6 := &List{6, nil}
	n1.next = n2
	n2.next = n5
	n5.next = n6
	n3.next = n4
	n4.next = n5
	fmt.Println(f(n1, n3))
}
