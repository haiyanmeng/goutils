package main

import (
	"fmt"

	"github.com/hmeng-19/goutils/stack"
)

func testStack() {
	s := stack.NewStack()
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
		s.Push(i + 20)
		s.Print()
	}
}

func testMaxStack() {
	s := stack.NewMaxStack()
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
	testStack()
	testMaxStack()
}
