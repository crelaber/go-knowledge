package linklist

//2.两数相加 中等
//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

//示例1
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.

//示例 2：
//输入：l1 = [0], l2 = [0]
//输出：[0]

//示例 3：
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]

func addTwoNum(l1 *ListNode, l2 *ListNode) *ListNode {
	var nums []int
	cur1, cur2 := l1, l2
	var carryBit bool //是否进位
	for cur1 != nil && cur2 != nil {
		sum := cur1.Val + cur2.Val
		if carryBit {
			sum++
		}
		carryBit = false
		if sum >= 10 {
			carryBit = true
		}

		nums = append(nums, sum%10) //取余
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	//计算两个listNode中不对称的N个数字
	nums = append(nums, traverse(cur1, carryBit)...)
	nums = append(nums, traverse(cur2, carryBit)...)
	if cur1 == nil && cur2 == nil && carryBit {
		nums = append(nums, 1)
	}
	return newList(nums)
}

func traverse(cur *ListNode, carryBit bool) (remainNums []int) {
	if cur == nil {
		return
	}

	for cur != nil {
		if carryBit { //有进位的运算，将结果加1
			res := cur.Val + 1
			if res >= 10 {
				remainNums = append(remainNums, 0)
				carryBit = true
			} else {
				remainNums = append(remainNums, res)
				carryBit = false
			}
			cur = cur.Next
			continue
		}
		remainNums = append(remainNums, cur.Val)
		cur = cur.Next
	}

	if carryBit {
		remainNums = append(remainNums, 1)
	}
	return
}
