func detectCapitalUse(word string) bool {
	/*
	   这题我学到了：
	   判断3种情况为true：
	   1 全大写
	   2 全小写
	   3 仅首字母大写
	   其余皆为false
	*/
	if strings.ToUpper(word[:1])+strings.ToLower(word[1:]) == word {
		return true
	}
	return strings.ToUpper(word) == word || strings.ToLower(word) == word
}