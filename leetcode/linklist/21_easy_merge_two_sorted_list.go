package linklist

//合并新链表
//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

//示例1
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]

//示例2
//输入：l1 = [], l2 = [0]
//输出：[0]

//提示：
//两个链表的节点数目范围是 [0, 50]
//-100 <= Node.val <= 100
//l1 和 l2 均按 非递减顺序 排列

func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := new(ListNode)
	dummy := cur
	cur1, cur2 := l1, l2

	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			cur.Next = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur2 = cur2.Next
		}
		cur = cur.Next
	}

	//处理剩余的链表，追加到新链表的尾部
	traverse := func(l, remainL *ListNode) {
		for remainL != nil {
			l.Next = remainL
			l = l.Next
			remainL = remainL.Next
		}
	}

	traverse(cur, cur1)
	traverse(cur, cur2)
	return dummy.Next
}
