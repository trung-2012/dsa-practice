func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxA := 0
	for left < right {
		if height[left] < height[right] {
			maxA = max(maxA, (right-left)*height[left])
			left++
		} else {
			maxA = max(maxA, (right-left)*height[right])
			right--
		}
	}
	return maxA
}