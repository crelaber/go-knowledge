package str

import "strings"

//最后一个单词的长度 简单
//给你一个字符串 s，由若干单词组成，单词之间用空格隔开。返回字符串中最后一个单词的长度。如果不存在最后一个单词，请返回 0 。

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	if len(words) <= 0 {
		return 0
	}
	return len(words[len(words)-1])
}
