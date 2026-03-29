func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}
	var keys []int
	for key := range count {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return count[keys[i]] > count[keys[j]]
	})
	return keys[:k]
}