func characterReplacement(s string, k int) int {
	count := make([]int, 26)
	i, res := 0, 0
	maxCount := 0

	for j := 0; j < len(s); j++ {
		count[s[j]-'A']++
		if count[s[j]-'A'] > maxCount {
			maxCount = count[s[j]-'A']
		}
		if j-i+1-maxCount > k {
			count[s[i]-'A']--
			i++
		}
		res = max(res, j-i+1)
	}
	return res
}