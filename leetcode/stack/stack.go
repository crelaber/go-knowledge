package stack

import (
	"errors"
	"log"
)

var ERR_EMPTY_STACK = errors.New("stack is empty")

type Stack []int

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(elem int) {
	*s = append(*s, elem)
}

func (s *Stack) Pop() int {
	n := s.Len()
	if n == 0 {
		log.Fatal(ERR_EMPTY_STACK)
	}
	peek := (*s)[n-1]
	*s = (*s)[:n-1]
	return peek
}

func (s *Stack) Peek() int {
	n := s.Len()
	if n == 0 {
		log.Fatal(ERR_EMPTY_STACK)
	}
	return (*s)[n-1]
}
