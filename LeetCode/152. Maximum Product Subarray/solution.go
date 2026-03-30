func maxProduct(nums []int) int {
	maxProd, minProd, res := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		temp := max(num, max(num*maxProd, num*minProd))
		minProd = min(num, min(num*maxProd, num*minProd))
		maxProd = temp

		res = max(res, maxProd)
	}
	return res
}