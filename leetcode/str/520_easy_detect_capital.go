package str

//检测大写字母
//给定一个单词，你需要判断单词的大写使用是否正确。
//我们定义，在以下情况时，单词的大写用法是正确的：
//全部字母都是大写，比如”USA”。
//单词中所有字母都不是大写，比如”leetcode”。
//如果单词不只含有一个字母，只有首字母大写， 比如 “Google”。
//否则，我们定义这个单词没有正确使用大写字母。

//示例 1:
//输入: "USA"
//输出: True

//示例 2:
//输入: "FlaG"
//输出: False

func detectCapitalUse(word string) bool {
	if len(word) <= 1 {
		return true
	}

	firstUp := rune(word[0]) >= 'A' && rune(word[0]) <= 'Z'
	allUpCase, allLowCase := true, true
	for i := 1; i < len(word); i++ {
		r := rune(word[i])
		if r < 'A' || r > 'Z' {
			allUpCase = false
		}

		if r < 'a' || r > 'z' {
			allLowCase = false
		}
	}

	if firstUp && (allLowCase || allUpCase) {
		return true
	}

	if !firstUp && allLowCase {
		return true
	}

	return false
}
