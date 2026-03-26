func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int

	var dfs func(start int, sum int)
	dfs = func(start int, sum int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		if sum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])

			dfs(i, sum+candidates[i])

			path = path[:len(path)-1]
		}
	}

	dfs(0, 0)
	return res
}