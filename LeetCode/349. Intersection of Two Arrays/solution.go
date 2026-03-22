func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	res := []int{}

	for _, v := range nums1 {
		m[v] = true
	}

	for _, v := range nums2 {
		if m[v] {
			res = append(res, v)
			delete(m, v)
		}
	}
	return res
}