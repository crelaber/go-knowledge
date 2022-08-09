package linklist

import (
	"fmt"
)

type LinkListNode struct {
	Val  int
	Next *LinkListNode
}

func ListToIntSlice(head *LinkListNode) []int {
	limit := 100
	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func IntArrToList(nums []int) *LinkListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &LinkListNode{}
	t := l
	for _, v := range nums {
		t.Next = &LinkListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

func (l *LinkListNode) GetNodeWith(val int) *LinkListNode {
	res := l
	for res != nil {
		if res.Val == val {
			break
		}
		res = res.Next
	}
	return res
}

func IntSliceToListWithCycle(nums []int, pos int) *LinkListNode {
	head := IntArrToList(nums)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}
