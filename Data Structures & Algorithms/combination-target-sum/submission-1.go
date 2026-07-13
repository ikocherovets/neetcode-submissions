func combinationSum(nums []int, target int) [][]int {
	res := [][]int{}

	var dfs func(i int, cur []int, total int)

	// Contract:
	// dfs(i, cur, total) знаходить всі комбінації,
	// використовуючи числа nums[i:],
	// які в сумі можуть досягнути target.
	//
	// cur - поточна побудована комбінація.
	// total - поточна сума елементів у cur.
	//
	// Все до i вже вирішено.
	//
	// Приклад:
	// nums = [2,3,6,7]
	// cur = [2]
	// total = 2
	//
	// dfs(1, cur, total)
	//
	// означає:
	// "Я вже вибрав 2.
	//  Тепер шукаю всі варіанти з [3,6,7],
	//  щоб дійти до target."

	dfs = func(i int, cur []int, total int) {

		// Base case:
		// Якщо поточна сума дорівнює target,
		// ми знайшли правильну комбінацію.
		if total == target {
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}

		// Якщо:
		// - закінчили всі числа
		// - або перевищили target
		//
		// ця гілка більше не може дати відповідь.
		if i >= len(nums) || total > target {
			return
		}


		// Варіант 1:
		// Беремо nums[i].
		//
		// i НЕ збільшуємо,
		// тому що одне число можна використовувати повторно.
		cur = append(cur, nums[i])
		dfs(i, cur, total+nums[i])


		// Backtracking:
		// Скасовуємо останній вибір,
		// щоб перевірити варіант без nums[i].
		cur = cur[:len(cur)-1]


		// Варіант 2:
		// Не беремо nums[i].
		//
		// Переходимо до наступного числа.
		dfs(i+1, cur, total)
	}

	dfs(0, []int{}, 0)

	return res
}