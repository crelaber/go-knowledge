package str

import "strings"

//统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
func countSegment(s string) int {
	return len(strings.Fields(s))
}
