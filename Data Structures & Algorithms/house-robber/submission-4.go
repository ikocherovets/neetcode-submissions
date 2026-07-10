func rob(nums []int) int {

	type cache map[int]int

	c := make(cache)

	// контракт:
	// dfs(i) = максимальна сума, яку можна отримати, починаючи з будинку i.
	//
	// На кожному будинку є два варіанти:
	// 1. Пропустити цей будинок -> dfs(i + 1)
	// 2. Взяти гроші з цього будинку -> nums[i] + dfs(i + 2)

	var dfs func(i int) int

	dfs = func(i int) int {
		// вийшли за межі масиву
		if i >= len(nums) {
			return 0
		}
		// якщо вже рахували цей індекс — повертаємо готову відповідь
		if val, ok := c[i]; ok {
			return val
		}
		// рахуємо найкращий варіант
		result := max(
			dfs(i+1),
			nums[i]+dfs(i+2),
		)
		// зберігаємо результат
		c[i] = result

		return result
	}

	return dfs(0)
}