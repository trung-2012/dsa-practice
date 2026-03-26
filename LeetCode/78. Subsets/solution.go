func subsets(nums []int) [][]int {
	var res [][]int
	var subset []int

	var backtrack func(start int)
	backtrack = func(start int) {
		temp := make([]int, len(subset))
		copy(temp, subset)
		res = append(res, temp)
		for i := start; i < len(nums); i++ {
			subset = append(subset, nums[i])
			backtrack(i + 1)
			subset = subset[:len(subset)-1]
		}
	}
	backtrack(0)
	return res
}