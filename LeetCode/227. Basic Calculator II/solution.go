func calculate(s string) int {
	curr := 0
	prev := 0
	res := 0
	op := '+'

	for i, ch := range s {
		if unicode.IsDigit(ch) {
			curr = curr*10 + int(ch-'0')
		}
		if !unicode.IsSpace(ch) && !unicode.IsDigit(ch) || i == len(s)-1 {
			switch op {
			case '+':
				res += prev
				prev = curr
			case '-':
				res += prev
				prev = -curr
			case '*':
				prev *= curr
			case '/':
				prev /= curr
			}
			op = ch
			curr = 0
		}
	}
	res += prev
	return res
}