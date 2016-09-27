package stack

import "fmt"

type GenericStack struct {
	data []interface{}
	n    int
}

func NewGenericStack() *GenericStack {
	return &GenericStack{}
}

func (s *GenericStack) Count() int {
	return s.n
}

func (s *GenericStack) Push(x interface{}) {
	s.data = append(s.data, x)
	s.n++
}

func (s *GenericStack) Pop() (interface{}, error) {
	if s.n == 0 {
		return 0, fmt.Errorf("the stack is empty")
	}

	x := s.data[s.n-1]
	s.data = s.data[:(s.n - 1)]
	s.n--
	return x, nil
}
