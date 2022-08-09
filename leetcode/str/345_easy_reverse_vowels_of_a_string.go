package str

//反转字符串中的元音字母 简单
//编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
//示例1
//输入："hello"
//输出："holle"

//示例2
//输入："leetcode"
//输出："leotcede"

var (
	vowelMap = map[byte]byte{
		'a': 'a',
		'A': 'A',
		'e': 'e',
		'E': 'E',
		'o': 'o',
		'O': 'O',
		'u': 'u',
		'U': 'U',
		'i': 'i',
		'I': 'I',
	}
)

func reverseVowels(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; {
		if !isVowel(b[i]) {
			i++
			continue
		}
		if !isVowel(b[j]) {
			j--
			continue
		}
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}

//判断是否为元音字母
func isVowel(s byte) bool {
	if _, ok := vowelMap[s]; ok {
		return true
	}
	return false
}
