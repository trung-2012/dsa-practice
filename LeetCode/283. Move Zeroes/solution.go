func moveZeroes(nums []int) {
	m := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if m != i {
				nums[m] = nums[i]
				nums[i] = 0
			}
			m++
		}

	}
}