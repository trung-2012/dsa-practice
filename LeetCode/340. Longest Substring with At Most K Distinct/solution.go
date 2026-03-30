func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 {
		return 0
	}

	count := make(map[byte]int)
	left := 0
	res := 0

	for right := 0; right < len(s); right++ {
		count[s[right]]++
		for len(count) > k {
			count[s[left]]--
			if count[s[left]] == 0 {
				delete(count, s[left])
			}
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}