package leetcode

import "testing"

func TestNewLRU(t *testing.T) {
	l := NewLRU(10)
	_ = l.Add("sa1", 1)
	_ = l.Add("sa2", 2)
	val, err := l.Get("sa1")
	if err {

	}
	print(val)
}
