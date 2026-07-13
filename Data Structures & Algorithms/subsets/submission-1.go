func subsets(nums []int) [][]int {
	res := [][]int{}
	subset := []int{}

	var dfs func(i int)

	// Contract:
	// dfs(i) добудовує всі можливі підмножини,
	// використовуючи елементи nums[i:] (від i до кінця).
	//
	// subset вже містить всі елементи,
	// які ми вирішили взяти до цього моменту.
	//
	// Все, що було до i — вже вирішено.
	// dfs(i) відповідає тільки за nums[i:]
	//
	// Наприклад:
	// nums = [1,2,3]
	// subset = [1]
	// dfs(1)
	//
	// означає:
	// "Я вже вибрав 1.
	//  Тепер добудую всі варіанти для [2,3]."

	dfs = func(i int) {
		// Base case:
		// Якщо всі елементи розглянули,
		// поточний subset є готовою відповіддю.
		if i >= len(nums) {
			temp := make([]int, len(subset))
			copy(temp, subset)
			res = append(res, temp)
			return
		}

		// Варіант 1: беремо nums[i]
		subset = append(subset, nums[i])
		dfs(i + 1)

		// Backtracking:
		// скасовуємо вибір, щоб спробувати
		// варіант, де nums[i] не беремо.
		subset = subset[:len(subset)-1]

		// Варіант 2: не беремо nums[i]
		dfs(i + 1)
	}

	dfs(0)

	return res
}