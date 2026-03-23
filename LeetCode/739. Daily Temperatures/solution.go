func dailyTemperatures(T []int) []int {
	n := len(T)
	res := make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return res
}