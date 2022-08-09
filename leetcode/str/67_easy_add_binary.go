package str

//二进制求和 简单
//给你两个二进制字符串，返回它们的和（用二进制表示）。

//示例1
//输入: a = "11", b = "1"
//输出: "100"

//示例2
//输入: a = "1010", b = "1011"
//输出: "10101"

//提示：
//每个字符串仅由字符 ‘0’ 或 ‘1’ 组成。
//1 <= a.length, b.length <= 10^4
//字符串如果不是 “0” ，就都不含前导零。

// 任意进制的任意长度数字相加，需从前向后遍历，求和后取余，并考虑进位
func addBinary(a string, b string) string {
	i1, i2 := len(a)-1, len(b)-1
	res := ""
	carry := 0 //初始的结果为0
	for i1 >= 0 || i2 >= 0 {
		sum := carry //取上次的结果看是否有仅为
		if i1 >= 0 {
			sum += int(rune(a[i1]) - '0')
			i1--
		}
		if i2 >= 0 {
			sum += int(rune(b[i2]) - '0')
			i2--
		}

		carry = sum / 2
		sum = sum % 2
		if sum == 0 {
			res += "0"
		} else {
			res += "1"
		}
	}

	if carry > 0 {
		res += "1"
	}
	return reverseStr(res)
}

func reverseStr(s string) string {
	runes := []rune(s)
	n := len(runes)
	mid := n / 2
	for i := 0; i < mid; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes)
}
