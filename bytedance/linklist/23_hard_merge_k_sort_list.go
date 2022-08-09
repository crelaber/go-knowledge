package linklist

//题目大意
//合并 K 个有序链表
//
//解题思路
//借助分治的思想，把 K 个有序链表两两合并即可。相当于是第 21 题的加强版。

func MergeTKLists(lists []*LinkListNode) *LinkListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}

	num := length / 2
	left := MergeTKLists(lists[:num])
	right := MergeTKLists(lists[num:])
	l := mergeTwoList(left, right)
	return l
}

func mergeTwoList(l1 *LinkListNode, l2 *LinkListNode) *LinkListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoList(l1.Next, l2)
		return l1
	}

	l2.Next = mergeTwoList(l1, l2.Next)
	return l2
}
