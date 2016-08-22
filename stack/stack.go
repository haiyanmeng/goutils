package stack

import "fmt"

type Stack struct {
	data []int
	n    int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Data() []int {
	return s.data
}

func (s *Stack) Count() int {
	return s.n
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

func (s *Stack) Peek() (int, error) {
	if s.n == 0 {
		return 0, fmt.Errorf("the stack is empty")
	}

	x := s.data[s.n-1]
	return x, nil
}

func (s *Stack) Print() {
	fmt.Printf("the stack is: %v\n", s.data)
}

type Elem struct {
	data int
	max  int
}

type MaxStack struct {
	data []Elem
	n    int
}

func NewMaxStack() *MaxStack {
	return &MaxStack{}
}

func (s *MaxStack) Push(x int) {
	var max int
	if s.n == 0 {
		max = x
	} else {
		max = s.data[s.n-1].max
	}
	if x > max {
		max = x
	}
	s.data = append(s.data, Elem{x, max})
	s.n++
}

func (s *MaxStack) Pop() (int, error) {
	if s.n == 0 {
		return 0, fmt.Errorf("the stack is empty")
	}
	x := s.data[s.n-1].data
	s.data = s.data[:(s.n - 1)]
	s.n--
	return x, nil
}

func (s *MaxStack) Max() (int, error) {
	if s.n == 0 {
		return 0, fmt.Errorf("the stack is empty")
	}
	return s.data[s.n-1].max, nil
}

func (s *MaxStack) Print() {
	for i := 0; i < s.n; i++ {
		fmt.Printf("%v(max:%v) ", s.data[i].data, s.data[i].max)
	}
	fmt.Printf("\n")
}
