package linklist

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	//l1 := newList([]int{8, 9, 9})
	//printListNode(l1)
	//l2 := newList([]int{2})
	//printListNode(l2)
	//cur := addTwoNum(l1, l2)
	//printListNode(cur)

	l3 := newList([]int{4, 5, 4, 1, 4})
	printListNode(l3)
	l4 := newList([]int{3, 5, 1, 5, 5, 4})
	printListNode(l4)
	cur2 := addTwoNum(l3, l4)
	printListNode(cur2)
}

func TestRemoveNthFromEnd(t *testing.T) {
	nums := []int{1, 2, 1, 34, 134}
	n := 3
	list := newList(nums)
	printListNode(list)
	targetNode := removeNthFromEnd(list, n)
	printListNode(targetNode)

}

//合并两个有序的链表
func TestMergeTwoSortedList(t *testing.T) {
	l1 := newList([]int{1, 2, 3, 5, 6})
	l2 := newList([]int{2, 2, 3, 4, 7, 9})
	node := mergeTwoList(l1, l2)
	printListNode(node)
}

func TestSwapPairsNode(t *testing.T) {
	l := newList([]int{1, 2, 3, 4, 5, 6, 7})
	printListNode(l)
	l = swapPairsNode(l)
	printListNode(l)
}

func printListNode(cur *ListNode) {
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
	fmt.Println("")
}
