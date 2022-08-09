package str

//实现 strStr() 函数。
//给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。
//说明：
//当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。

//示例1
//输入：haystack = "hello", needle = "ll"
//输出：2

//示例2
//输入：haystack = "aaaaa", needle = "bba"
//输出：-1

//示例三
//输入：haystack = "", needle = ""
//输出：0

func strStr(hack string, needle string) int {
	if len(needle) <= 0 {
		return 0
	}

	if len(needle) > len(hack) {
		return -1
	}

	hs, ns := []rune(hack), []rune(needle)
	lh, ln := len(hs), len(ns)
	var starts []int
	for i, h := range hs {
		if h == ns[0] {
			starts = append(starts, i)
		}
	}

	for _, start := range starts {
		if start+ln > lh {
			break
		}
		if string(hs[start:start+ln]) == needle {
			return start
		}
	}
	return -1
}
