package linklist

//题目：回文链表
//题目大意 #
//判断一个链表是否是回文链表。要求时间复杂度 O(n)，空间复杂度 O(1)。

//解题思路 #
//这道题只需要在第 143 题上面改改就可以了。思路是完全一致的。先找到中间结点，然后反转中间结点后面到结尾的所有结点。最后一一判断头结点开始的结点和中间结点往后开始的结点是否相等。如果一直相等，就是回文链表，如果有不相等的，直接返回不是回文链表

// IsPalindrome 解法1
func IsPalindrome(head *LinkListNode) bool {
	slice := []int{}
	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}

	for i, j := 0, len(slice)-1; i < j; {
		if slice[i] != slice[j] {
			return false
		}
		i++
		j--
	}
	return true
}

//解法2
func isPalindrome(head *LinkListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	res := true
	p1 := head
	p2 := head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	//翻转链表后半部分
	preMiddle := p1
	preCurrent := p1.Next
	for preCurrent != nil {
		current := preCurrent.Next
		preCurrent.Next = current.Next
		preMiddle.Next = current
	}
	//扫描链表判断是否为回文数
	p1 = head
	p2 = preMiddle.Next
	for p1 != preMiddle {
		if p1.Val == p2.Val {
			p1 = p1.Next
			p2 = p2.Next
		} else {
			res = false
			break
		}
	}

	if p1 == preMiddle {
		if p2 != nil && p1.Val != p2.Val {
			return false
		}
	}
	return res
}
