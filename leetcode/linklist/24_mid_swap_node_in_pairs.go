package linklist

//24. 两两交换链表中的节点 中等
//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

//示例1
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
//示例 2：
//输入：head = []
//输出：[]
//
//示例 3：
//输入：head = [1]
//输出：[1]

func swapPairsNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { // case [] 或者 [1]
		return head
	}
	newHead := swap(head, head.Next)
	pre := head
	if pre.Next == nil || pre.Next.Next == nil { //case [1,2]
		return newHead
	}

	cur := pre.Next
	next := cur.Next
	for cur != nil && next != nil {
		pre.Next = swap(cur, next)
		if cur.Next == nil {
			break
		}
		pre = cur
		cur = cur.Next
		next = cur.Next
	}
	return newHead
}

func swap(cur, next *ListNode) *ListNode {
	cur.Next = next.Next
	next.Next = cur
	return next
}
