func toInt(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func evalRPN(tokens []string) int {
	stack := make([]string, len(tokens))
	index := -1
	for _, t := range tokens {
		if strings.Index("+-*/", t) >= 0 {
			index--
			stack[index] = compute(stack[index], stack[index+1], t)
			continue
		}
		index++
		stack[index] = t
	}
	return toInt(stack[0])
}

var toStr = strconv.Itoa

func compute(x, y, t string) string {
	var res string
	if t == "+" {
		res = toStr(toInt(x) + toInt(y))
	} else if t == "-" {
		res = toStr(toInt(x) - toInt(y))
	} else if t == "*" {
		res = toStr(toInt(x) * toInt(y))
	} else if t == "/" {
		res = toStr(toInt(x) / toInt(y))
	}
	return res
}