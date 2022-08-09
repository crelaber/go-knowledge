package str

import "strings"

//最长公共前缀 简单
//编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 “”。
//示例1
//输入：strs = ["flower","flow","flight"]
//输出："fl"

//示例2
//输入：strstrs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。

//提示：
//0 <= strs.length <= 200
//0 <= strs[i].length <= 200
//strs[i] 仅由小写英文字母组成

//维护一个前缀库，不断往后遍历，判断前缀逐步剪短前缀库的大小
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	s1 := strs[0]
	prefixes := make([]string, len(s1))
	for i := range s1 {
		prefixes[i] = s1[:i+1]
	}

	for _, s := range strs {
		for i, pre := range prefixes {
			if !strings.HasPrefix(s, pre) {
				prefixes = prefixes[:i]
				break
			}
		}
	}

	if len(prefixes) > 0 {
		return prefixes[len(prefixes)-1]
	}
	return ""
}
