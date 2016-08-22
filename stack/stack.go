package stack

import "fmt"

type Stack struct {
	data []int
	n    int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
	s.n++
}

func (s *Stack) Pop() (int, error) {
	if s.n == 0 {
		return 0, fmt.Errorf("the stack is empty")
	}

	x := s.data[s.n-1]
	s.data = s.data[:(s.n - 1)]
	s.n--
	return x, nil
}

func (s *Stack) Print() {
	fmt.Printf("the stack is: %v\n", s.data)
}
