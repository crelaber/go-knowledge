package string

import "go-knowledge/leetcode/utils"

//题目大意
//给你一个字符串 s，找到 s 中最长的回文子串。
//实现参考：https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0005.Longest-Palindromic-Substring/

//解法一，动态规划。定义 dp[i][j] 表示从字符串第 i 个字符到第 j 个字符这一段子串是否是回文串。
//由回文串的性质可以得知，回文串去掉一头一尾相同的字符以后，剩下的还是回文串。
//所以状态转移方程是 dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])，
//注意特殊的情况，j - i == 1 的时候，即只有 2 个字符的情况，只需要判断这 2 个字符是否相同即可。
//j - i == 2 的时候，即只有 3 个字符的情况，只需要判断除去中心以外对称的 2 个字符是否相等。
//每次循环动态维护保存最长回文串即可。时间复杂度 O(n^2)，空间复杂度 O(n^2)。

//解法一，Manacher算法，时间复杂度o(n)，空间复杂度o(n)
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	newS := make([]rune, 0)
	newS = append(newS, '#')
	//将字符串变为的长度变为基数个
	for _, c := range s {
		newS = append(newS, c)
		newS = append(newS, '#')
	}

	// dp[i]:    以预处理字符串下标 i 为中心的回文半径(奇数长度时不包括中心)
	// maxRight: 通过中心扩散的方式能够扩散的最右边的下标
	// center:   与 maxRight 对应的中心字符的下标
	// maxLen:   记录最长回文串的半径
	// begin:    记录最长回文串在起始串 s 中的起始下标

	dp, maxRight, center, maxLen, begin := make([]int, len(newS)), 0, 0, 1, 0
	for i := 0; i < len(newS); i++ {
		if i < maxRight {
			dp[i] = utils.Min(maxRight-i, dp[2*center-i]) //这一步是Manacher算法的关键所在
		}
		//中心扩散法更新dp[i]
		left, right := i-(1+dp[i]), i+(1+dp[i])
		for left >= 0 && right <= len(newS) && newS[left] == newS[right] {
			dp[i]++
			left--
			right++
		}
		// 更新 maxRight，它是遍历过的 i 的 i + dp[i] 的最大者
		if i+dp[i] > maxRight {
			maxRight = i + dp[i]
			center = i
		}

		// 记录最长回文子串的长度和相应它在原始字符串中的起点

		if dp[i] > maxLen {
			maxLen = dp[i]
			begin = (i - maxLen) / 2 // 这里要除以 2 因为有我们插入的辅助字符 #
		}
	}
	return s[begin : begin+maxLen]
}
