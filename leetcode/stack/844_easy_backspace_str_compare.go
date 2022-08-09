package stack

import "strings"

//题目：比较含退格的字符串
//给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。
//注意：如果对空文本输入退格字符，文本继续为空。

//示例 1：
//输入：S = "ab#c", T = "ad#c"
//输出：true
//解释：S 和 T 都会变成 “ac”。

//示例 2：
//输入：S = "ab##", T = "c#d#"
//输出：true
//解释：S 和 T 都会变成 “”。

//示例 3：
//输入：S = "a##c", T = "#a#c"
//输出：true
//解释：S 和 T 都会变成 “c”。

//示例 4：
//输入：S = "a#c", T = "b"
//输出：false
//解释：S 会变成 “c”，但 T 仍然是 “b”。

func backspaceCompare(S string, T string) bool {
	var s1, s2 Stack
	travese := func(str string, s *Stack) {
		for _, r := range str {
			if r == '#' {
				if !s.IsEmpty() {
					s.Pop()
				}
			} else {
				s.Push(int(r))
			}
		}
	}
	travese(S, &s1)
	travese(T, &s2)
	var s, t string
	for !s1.IsEmpty() {
		s += string(rune(s1.Pop()))
	}
	for !s2.IsEmpty() {
		t += string(rune(s2.Pop()))
	}
	return strings.EqualFold(s, t)
}
