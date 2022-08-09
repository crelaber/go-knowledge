package str

// ReverseString 翻转一个字符串
//解题思路：使用两个指针对撞，不断的交换首尾元素
func ReverseString(s []byte) string {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return string(s)
}
