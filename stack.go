package tdd_stack

import (
	"fmt"
)

type Stack struct {
	arr []int
}

var ErrNoSuchElement = fmt.Errorf("no such element")

func New() *Stack {
	return &Stack{
		arr: make([]int, 0),
	}
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.arr)
}

func (s *Stack) Push(val int) error {
	s.arr = append(s.arr, val)
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrNoSuchElement
	}
	val := s.getLastElement()
	s.arr = s.arr[:s.Size() - 1]
	return val, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, ErrNoSuchElement
	}
	return s.getLastElement(), nil
}

func (s *Stack) getLastElement() int {
	return s.arr[s.Size() - 1]
}