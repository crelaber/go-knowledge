package stack

//题目：856. 括号的分数 中等
//给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：
//() 得 1 分。
//AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
//(A) 得 2 * A 分，其中 A 是平衡括号字符串。

//示例 1：
//输入： "()"
//输出： 1
//
//示例 2：
//输入： "(())"
//输出： 2
//
//示例 3：
//输入： "()()"
//输出： 2
//
//示例 4：
//输入： "(()(()))"
//输出： 6

func scoreOfParentheses(S string) int {
	var s Stack
	for _, r := range S {
		switch r {
		case '(':
			s.Push(-1)
		case ')':
			cur := 0
			for s.Peek() != -1 {
				cur += s.Pop()
			}
			s.Pop()
			if cur == 0 {
				s.Push(1)
			} else {
				s.Push(cur * 2)
			}
		}
	}

	sum := 0
	for !s.IsEmpty() {
		sum += s.Pop()
	}
	return sum
}
