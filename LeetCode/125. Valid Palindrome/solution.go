func isPalindrome(s string) bool {
	result := []rune{}
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			result = append(result, v)
		} else if v >= 'A' && v <= 'Z' {
			result = append(result, v+32)
		} else if v >= '0' && v <= '9' {
			result = append(result, v)
		}
	}
	for i := 0; i < len(result)/2; i++ {
		if result[i] != result[len(result)-i-1] {
			return false
		}
	}
	return true
}