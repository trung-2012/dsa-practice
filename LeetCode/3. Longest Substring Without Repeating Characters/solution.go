func lengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int)
	i, res := 0, 0
	for j, v := range s {
		if idx, found := seen[v]; found {
			i = max(i, idx+1)
		}
		seen[v] = j
		res = max(res, j-i+1)
	}
	return res
}