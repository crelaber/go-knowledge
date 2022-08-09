package string

import "sort"

//编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 “"。

//解题思路 #
//对 strs 按照字符串长度进行升序排序，求出 strs 中长度最小字符串的长度 minLen
//逐个比较长度最小字符串与其它字符串中的字符，如果不相等就返回 commonPrefix,否则就把该字符加入 commonPrefix

func longestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	minLen := len(strs[0])
	if minLen == 0 {
		return ""
	}

	var commonPrefix []byte
	for i := 0; i < minLen; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] {
				return string(commonPrefix)
			}
		}
		commonPrefix = append(commonPrefix, strs[0][i])
	}
	return string(commonPrefix)
}
