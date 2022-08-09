package str

import (
	"fmt"
	"testing"
)

//最长回文数算法
func TestLongestPalindrome(t *testing.T) {
	s := "ababeadkadidakwaiewaewew"
	p := bestLongestPalindrome(s)
	fmt.Println("p===>" + p)

}

//翻转字符串
func TestReverseStr(t *testing.T) {
	s := "a311-29383sldf"
	b := ReverseString([]byte(s))
	str := string(b)
	fmt.Println("str===>" + str)
}

//
func TestZigZagConvert(t *testing.T) {
	s := "PAYPALISHIRING"
	fmt.Println("CONVERSION:" + convertZigZag(s, 4))
}

//将字符串转化为整数，有10，-10，+10三种情况
func TestMyAtoi(t *testing.T) {
	str := "-13838"
	fmt.Printf("%d", myAtoi(str))
}

//将整数转化为罗马数字
func TestIntToRoman(t *testing.T) {
	num := 1023
	roman := intToRoman(num)
	fmt.Println(roman)
}

//最长公共前缀
func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"common", "company", "cooperation"}
	str := longestCommonPrefix(strs)
	fmt.Println(str)
}

//外观队列，后一个字符是前一个字符的描述
func TestCountAndSay(t *testing.T) {
	for i := 1; i <= 15; i++ {
		str := countAndSay(i)
		fmt.Println(str)
	}
}

//最后一个字符的长度，使用strings.Fields方法获取各个words
func TestLengthOfLastWords(t *testing.T) {
	s := "hello world you a get a new way"
	l := lengthOfLastWord(s)
	fmt.Println(l)
}

//二进制的加法
func TestAddBinary(t *testing.T) {
	a := "10110"
	b := "1011110"
	fmt.Println(addBinary(a, b))
}

//判断是否为为回文数
func TestIsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}

//第一个不重复的字符
func TestFirstUniqChar(t *testing.T) {
	s := "leelcode"
	fmt.Println(firstUniqChar(s))
}

//检测单词是否正确的使用了大写
func TestDetectCapital(t *testing.T) {
	s := "flag"
	fmt.Println(detectCapitalUse(s))
	s = "flaG"
	fmt.Println(detectCapitalUse(s))
	s = "Flag"
	fmt.Println(detectCapitalUse(s))
}
