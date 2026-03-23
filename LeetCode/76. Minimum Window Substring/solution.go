func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	need := make(map[byte]int)
	for j := 0; j < len(t); j++ {
		need[t[j]]++
	}
	window := make(map[byte]int)

	have, needCount := 0, len(need)
	res := ""
	resLen := len(s) + 1

	i := 0

	for j := 0; j < len(s); j++ {
		window[s[j]]++
		if need[s[j]] > 0 && window[s[j]] == need[s[j]] {
			have++
		}
		for have == needCount && i <= j {
			if j-i+1 < resLen {
				res = s[i : j+1]
				resLen = j - i + 1
			}

			left := s[i]
			window[left]--
			if need[left] > 0 && need[left] > window[left] {
				have--
			}
			i++
		}
	}
	return res
}