package stack

//给定一个经过编码的字符串，返回它解码后的字符串。
//编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
//示例1
//输入：s = "3[a]2[bc]"
//输出："aaabcbc"
//
//示例 2：
//输入：s = "3[a2[c]]"
//输出："accaccacc"
//
//示例 3：
//输入：s = "2[abc]3[cd]ef"
//输出："abcabccdcdcdef"

func decodeString(s string) string {
	var (
		n    int
		res  string
		nums Stack
		str  StrStack
	)
	for _, r := range s {
		switch {
		case '0' <= r && r <= '9':
			n = 10*n + int(r-'0')
		case ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z'):
			res += string(r)
		case r == '[':
			nums.Push(n)
			n = 0
			str.Push(res)
			res = ""
		case r == ']':
			repeat := res
			for i := 0; i < nums.Peek()-1; i++ {
				res += repeat
			}
			res = str.Pop() + res
			nums.Pop()
		}
	}
	return res
}
