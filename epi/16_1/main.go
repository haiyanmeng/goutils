package main

import (
	"fmt"

	"github.com/hmeng-19/goutils/stack"
)

func movePeg(A, B, C *stack.Stack, n int) {
	if n == 0 {
		return
	}

	if n == 1 {
		x, _ := A.Pop()
		B.Push(x)
		fmt.Printf("move %v from %v to %v\n", x, A.Name, B.Name)
		return
	}

	movePeg(A, C, B, n-1)
	movePeg(A, B, C, 1)
	movePeg(C, B, A, n-1)
}

func main() {
	s := stack.NewStack()
	s.Name = "A"
	for i := 3; i > 0; i-- {
		s.Push(i)
	}
	b, c := stack.NewStack(), stack.NewStack()
	b.Name, c.Name = "B", "C"
	s.Print()
	b.Print()
	c.Print()
	movePeg(s, b, c, s.Count())
}
