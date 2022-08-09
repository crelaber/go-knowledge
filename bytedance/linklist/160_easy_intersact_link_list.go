package linklist

import "fmt"

//
//题目：找到两个链表的交叉点
//
//题目大意 #
//找到 2 个链表的交叉点。
//
//解题思路 #
//这道题的思路其实类似链表找环。
//
//给定的 2 个链表的长度如果一样长，都从头往后扫即可。如果不一样长，需要先“拼成”一样长。
//把 B 拼接到 A 后面，把 A 拼接到 B 后面。这样 2 个链表的长度都是 A + B。再依次扫描比较 2 个链表的结点是否相同。

func getIntersectionNode(headA, headB *LinkListNode) *LinkListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
		fmt.Printf("a = %v b=%v\n", a, b)
	}
	return a
}
