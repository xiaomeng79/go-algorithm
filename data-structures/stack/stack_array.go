package stack

import (
	"errors"
	"sync"
)

const ARRAY_SIZE = 10

//数组实现
type Stack struct {
	data [ARRAY_SIZE]int
	top  int
	lock sync.Mutex
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.top
}

func (s *Stack) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.top == 0
}

func (s *Stack) Push(i int) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.top == ARRAY_SIZE {
		return errors.New("栈已满")
	}
	s.data[s.top] = i
	s.top++
	return nil
}
func (s *Stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.top == 0 {
		return 0, errors.New("栈空")
	}
	s.top--
	return s.data[s.top], nil
}

func (s *Stack) Peek() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.top == 0 {
		return 0, errors.New("栈空")
	}
	return s.data[s.top-1], nil
}
