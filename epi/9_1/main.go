package main

import (
	"fmt"

	"github.com/hmeng-19/goutils/stack"
)

type MaxStack struct {
	dataStack stack.Stack
	maxStack  stack.Stack
}

func NewMaxStack() *MaxStack {
	return &MaxStack{}
}

func (s *MaxStack) Push(x int) {
	var max int
	top, err := s.maxStack.Peek()
	if err == nil {
		max = top
	} else {
		max = x
	}

	if x > max {
		max = x
	}
	s.dataStack.Push(x)
	s.maxStack.Push(max)
}

func (s *MaxStack) Pop() (int, error) {
	x, err := s.dataStack.Pop()
	if err != nil {
		return x, err
	}
	s.maxStack.Pop()
	return x, nil
}

func (s *MaxStack) Max() (int, error) {
	top, err := s.maxStack.Peek()
	if err != nil {
		return 0, err
	}

	return top, nil
}

func (s *MaxStack) Print() {
	n := s.dataStack.Count()
	for i := 0; i < n; i++ {
		fmt.Printf("%v(max:%v) ", s.dataStack.Data()[i], s.maxStack.Data()[i])
	}
	fmt.Printf("\n")
}

func testMaxStack() {
	s := NewMaxStack()
	for i := 0; i < 10; i++ {
		s.Push(i)
		s.Print()
	}
	for i := 0; i < 5; i++ {
		x, err := s.Pop()
		fmt.Printf("the popped item is %v\n", x)
		s.Print()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for i := 0; i < 10; i++ {
		s.Push(30 - i)
		s.Print()
		max, err := s.Max()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("the max is %v\n", max)
	}
}

func main() {
	testMaxStack()
}
