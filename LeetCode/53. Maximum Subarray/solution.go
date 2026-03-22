func maxSubArray(nums []int) int {
	sum, maxS := nums[0], nums[0]
	for _, v := range nums[1:] {
		if sum < 0 {
			sum = v
		} else {
			sum += v
		}
		if maxS < sum {
			maxS = sum
		}
	}
	return maxS
}