package main

func validPalindrome(s string) bool {
	/*
	   这题我学到了：
	       利用一次递归求解
	*/
	return helper(s, 0, len(s)-1, 1)
}

func helper(s string, i, j, n int) bool {
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			if n > 0 {
				return helper(s, i, j-1, 0) || helper(s, i+1, j, 0)
			}
			return false
		}
	}
	return true
}
