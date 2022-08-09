package str

import "go-knowledge/leetcode/utils"

//最长回文子串 中等
//最有解法，马车拉算法，检测最长回文子串
//时间复杂度o(n)
func bestLongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	t := "#"
	for _, r := range s {
		t += string(r) + "#" //统一将奇数偶数串 -> 奇数串
	}

	radius := make([]int, len(t)) //每个位置的回文半径
	maxCenter := 0                //到目前位置最长回文串的中心位置
	maxR := 0                     //到目前位置最长回文串的右边界索引
	stopCenter := 0
	maxRadius := 0

	for i := 0; i < len(t); i++ {
		mirror := 2*maxCenter - i
		if i < maxR {
			radius[i] = utils.Min(radius[mirror], maxR-i) //去【i，mx】与i半径的最小值
		}
		//以i为中心不断的扩充半径，向左右两边探测回文
		for i+1+radius[i] < len(t) && i-1-radius[i] >= 0 && t[i-1-radius[i]] == t[i+1+radius[i]] {
			radius[i]++
		}
		//超出了边界
		if i+radius[i] > maxRadius {
			maxR = i + radius[i]
			maxCenter = i
		}
		if radius[i] > maxRadius {
			maxRadius = radius[i]
			stopCenter = i
		}
	}
	l := stopCenter/2 - maxRadius/2 //t->s， 索引取中
	r := l + maxRadius - 1          //减掉自身
	return string(s[l : r+1])
}
