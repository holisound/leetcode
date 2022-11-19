func buddyStrings(A string, B string) bool {
	if len(a) != len(b) { // 长度不等，肯定为false
		return false
	}
	count := map[rune]int{}
	dim := 0         //字符出现最高的频次
	diff := 0        //A和B，同索引异字符的总数
	p1, p2 := -1, -1 //前2次出现异字符的索引
	for i, a := range A {
		count[a]++
		if count[a] > dim {
			dim = count[a]
		}
		if A[i] != B[i] {
			diff++
			if diff == 1 {
				p1 = i
			} else if diff == 2 {
				p2 = i
			} else {
				return false
			}
		}
	}
	if A == B && dim > 1 {
		return true
	}
	if p1 != -1 && p2 != -1 && A[p1] == B[p2] && A[p2] == B[p1] {
		return true
	}
	return false

}