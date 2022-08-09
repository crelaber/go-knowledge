package linklist

//19.删除链表的倒数第 N 个结点 中等
//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//进阶：你能尝试使用一趟扫描实现吗？

//示例1
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]

//示例 2：
//输入：head = [1], n = 1
//输出：[]

//示例 3：
//输入：head = [1,2], n = 1
//输出：[1]

// 和倒数相关的问题考虑天生有间距的"双指针"，一次遍历解决
// 借助哑节点是关键，用于防止删除倒数第 len(nums) 个节点会导致链表丢失
// 若要兼顾正常节点的移动，又要处理头结点的特殊情况，请考虑"哑节点"
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Val: 0, Next: head}
	front, rear := dummyNode, dummyNode
	for counts := 0; counts <= n; counts++ {
		rear = rear.Next
	}
	for rear != nil {
		front = front.Next
		rear = rear.Next
	}
	front.Next = front.Next.Next //删除节点
	return dummyNode.Next
}
