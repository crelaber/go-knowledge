package linklist

//题目：翻转链表

func reverseLinkList(head *LinkListNode) *LinkListNode {
	var behind *LinkListNode
	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next
	}
	return behind
}
