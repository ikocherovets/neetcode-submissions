func combinationSum(nums []int, target int) [][]int {
    result := [][]int{}
	// contract -> повертає всі можливі комбінації для nums[i:] щоб сума == target. Тому треба 2 змінна - currentSlice and total
	var dfs func(i int, cur []int, total int) 
	dfs = func(i int, cur []int, total int) {
		if total == target {
			temp := make([]int, len(cur))
			copy(temp, cur)
			result = append(result, temp)
			return
		}

		if total > target || i >= len(nums) {
			return
		}

		// не додаю індекс оновлюю total бо current стає тим що в nums[i]
		cur = append(cur, nums[i])
		dfs(i, cur, total + nums[i])

		// backtracking - тепер якщо рухаюся далі
		// це все випливає з контракту чому така поведінка
		cur = cur[:len(cur) - 1]
		dfs(i + 1, cur, total)
	}


	dfs(0, []int{}, 0)

	return result
}
