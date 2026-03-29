import "math/rand"

func findKthLargest(nums []int, k int) int {
	target := len(nums) - k
	return quickSelect(nums, 0, len(nums)-1, target)
}

func quickSelect(nums []int, left int, right int, k int) int {
	if left == right {
		return nums[left]
	}
	pivotIndex := partition(nums, left, right)
	if pivotIndex == k {
		return nums[pivotIndex]
	} else if pivotIndex > k {
		return quickSelect(nums, left, pivotIndex-1, k)
	} else {
		return quickSelect(nums, pivotIndex+1, right, k)
	}
}

func partition(nums []int, left int, right int) int {
	pivotIndex := left + rand.Intn(right-left+1)
	pivot := nums[pivotIndex]

	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]

	storeIndex := left
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[i], nums[storeIndex] = nums[storeIndex], nums[i]
			storeIndex++
		}
	}
	nums[storeIndex], nums[right] = nums[right], nums[storeIndex]
	return storeIndex
}