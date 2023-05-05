package utils

import (
	"container/list"
	"sync"
)

type Stack struct {
	dll   *list.List
	mutex sync.Mutex
	count int
}

func NewStack() *Stack {
	return &Stack{dll: list.New(), count: 0}
}

func (s *Stack) Push(x interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.dll.PushBack(x)
	s.count++
}

func (s *Stack) Pop() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.dll.Len() == 0 {
		return nil
	}
	tail := s.dll.Back()
	val := tail.Value
	s.dll.Remove(tail)
	s.count--
	return val
}

func (s *Stack) Count() int {
	return s.count
}
