func subsets(nums []int) [][]int {
	res := [][]int{}
	subset := []int{}

	var dfs func(i int)
	dfs = func(i int) {
		if i >= len(nums) {
			cp := make([]int, len(subset))
			copy(cp, subset)
			res = append(res, cp)
			return
		}

		// decision to include nums[i]
		subset = append(subset, nums[i])
		dfs(i + 1)

		// decision NOT to include nums[i]
		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}

	dfs(0)
	return res
}