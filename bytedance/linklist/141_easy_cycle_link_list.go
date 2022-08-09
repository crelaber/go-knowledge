package linklist

//题目：判断链表是否有环，不能使用额外的空间
//题目大意 #=
//判断链表是否有环，不能使用额外的空间。
//解题思路 #
//给 2 个指针，一个指针是另外一个指针的下一个指针。快指针一次走 2 格，慢指针一次走 1 格。如果存在环，那么前一个指针一定会经过若干圈之后追上慢的指针。

func hasCycle(head *LinkListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next //快指针走两步
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
