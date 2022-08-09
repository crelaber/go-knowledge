package stack

import (
	"errors"
	"log"
)

var ERR_EMPTY_STR_STACK = errors.New("statck is empty")

type StrStack []string

func (s *StrStack) Len() int {
	return len(*s)
}

func (s *StrStack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *StrStack) Push(elem string) {
	*s = append(*s, elem)
}

func (s *StrStack) Pop() string {
	n := s.Len()
	if n == 0 {
		log.Fatalln(ERR_EMPTY_STR_STACK)
	}

	peek := (*s)[n-1]
	*s = (*s)[:n-1]
	return peek
}

func (s *StrStack) Peek() string {
	n := s.Len()
	return (*s)[n-1]
}
