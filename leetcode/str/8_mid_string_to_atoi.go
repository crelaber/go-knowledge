package str

import "strings"

//字符串转换整数 (atoi) 中等
//请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。
//函数 myAtoi(string s) 的算法如下
//1、读入字符串并丢弃无用的前导空格
//2、检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
//3、读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
//4、将前面步骤读入的这些数字转换为整数（即，”123” -> 123， “0032” -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
//5、如果整数数超过 32 位有符号整数范围 [−2的31次方, 2的31次方 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −2的31次方 的整数应该被固定为 −2的31次方 ，大于 2的31次方 − 1 的整数应该被固定为 2的31次方− 1 。
//6、返回整数作为最终结果。

const (
	INT_MIN = -1 << 31  //最小数
	INT_MAX = 1<<31 - 1 //最大数
)

//正常整数字符串只有3中case
//10
//+10
//-10
//需要注意溢出的情况
func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}

	positive := true //是否为正数的标志
	switch str[0] {
	case '-':
		positive = false
		str = str[1:]
	case '+':
		str = str[1:]
	}

	var nums []rune
	for _, r := range str { //将输入的string归并到数组中
		if r < '0' || r > '9' {
			break
		}
		nums = append(nums, r)
	}
	var n int
	for _, num := range nums {
		n = 10*n + int(num-'0')
		if positive && n > INT_MAX {
			return INT_MAX
		}

		if !positive && -n < INT_MIN {
			return INT_MIN
		}
	}

	if !positive {
		return -n
	}
	return n
}
