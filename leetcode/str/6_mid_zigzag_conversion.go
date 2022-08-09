package str

import "fmt"

//将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//比如输入字符串为 “PAYPALISHIRING” 行数为 3 时，排列如下：
//P   A   H   N
//A P L S I I G
//Y   I   R
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如：”PAHNAPLSIIGYIR”。

func convertZigZag(s string, numRows int) string {
	n := len(s)
	if n <= 1 || numRows <= 1 {
		return s
	}

	lines := make([][]rune, numRows)
	for i := 0; i < n; i++ {
		//向下走
		for r := 0; r <= numRows-1 && i < n; r++ {
			lines[r] = append(lines[r], rune(s[i]))
			i++
		}

		//向右上方走
		for r := numRows - 2; r >= 1 && i < n; r-- {
			lines[r] = append(lines[r], rune(s[i]))
			i++
		}
	}

	fmt.Println(lines)

	var str string
	for _, line := range lines {
		for _, r := range line {
			str += string(r)
		}
	}
	return str
}
