package main

import (
	"fmt"

	"github.com/hmeng-19/goutils/stack"
)

func main() {
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
