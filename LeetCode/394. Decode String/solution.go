type stackItem struct {
	str string
	num int
}

func decodeString(s string) string {
	stack := []stackItem{}
	current_num := 0
	current_string := ""

	for _, char := range s {
		if char >= '0' && char <= '9' {
			current_num = current_num*10 + int(char-'0')
		} else if char == '[' {
			stack = append(stack, stackItem{str: current_string, num: current_num})
			current_string = ""
			current_num = 0
		} else if char == ']' {
			lastIndex := len(stack) - 1
			poppedItem := stack[lastIndex]
			stack = stack[:lastIndex]
			current_string = poppedItem.str + strings.Repeat(current_string, poppedItem.num)
		} else {
			current_string += string(char)
		}
	}
	return current_string
}