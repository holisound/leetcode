import lls "github.com/emirpasic/gods/stacks/linkedliststack"

func decodeString(s string) string {
	stack := lls.New()
	var chars []rune
	for _, c := range s {
		if c == ']' {
			for val, ok := stack.Peek(); ok && unicode.IsLetter(val.(rune)); val, ok = stack.Peek() {
				chars = append(chars, val.(rune))
				stack.Pop()
			}
			reverse(chars)

			seed := chars
			chars = []rune{}
			stack.Pop()

			for val, ok := stack.Peek(); ok && unicode.IsDigit(val.(rune)); val, ok = stack.Peek() {
				chars = append(chars, val.(rune))
				stack.Pop()
			}
			reverse(chars)
			num, _ := strconv.Atoi(string(chars))

			for i := 0; i < num; i++ {
				for _, c := range seed {
					stack.Push(c)
				}
			}
			chars = []rune{}
			continue
		}

		stack.Push(c)

	}
	chars = []rune{}
	for val, ok := stack.Peek(); ok; val, ok = stack.Peek() {
		chars = append(chars, val.(rune))
		stack.Pop()
	}
	return string(reverse(chars))
}

func reverse(chars []rune) []rune {
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return chars
}
