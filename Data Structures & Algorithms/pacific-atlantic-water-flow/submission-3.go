func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])

	// множини клітинок, з яких вода може дотекти до кожного з океанів
	pac := make(map[[2]int]bool)
	atl := make(map[[2]int]bool)

	// оголошуємо тип наперед — інакше Go не дозволить рекурсивний виклик
	// всередині ще не визначеної функції
	var dfs func(r, c int, visit map[[2]int]bool, prevHeight int)

	dfs = func(r, c int, visit map[[2]int]bool, prevHeight int) {
		// стоп-умови:
		// - вийшли за межі сітки
		// - клітинку вже відвідали в цьому обході
		// - висота клітинки менша за попередню — вода "вгору" не тече,
		//   значить з поточної клітинки сюди дотекти не могло б
		if r < 0 || c < 0 || r == rows || c == cols ||
			visit[[2]int{r, c}] || heights[r][c] < prevHeight {
			return
		}

		// позначаємо клітинку як досяжну для цього океану
		visit[[2]int{r, c}] = true

		// йдемо у всі 4 сторони, "піднімаючись" по висоті від поточної клітинки
		dfs(r+1, c, visit, heights[r][c])
		dfs(r-1, c, visit, heights[r][c])
		dfs(r, c+1, visit, heights[r][c])
		dfs(r, c-1, visit, heights[r][c])
	}

	// стартуємо DFS від країв, що межують з Pacific (верхній ряд + лівий стовпець)
	// і від країв, що межують з Atlantic (нижній ряд + правий стовпець)
	for c := 0; c < cols; c++ {
		dfs(0, c, pac, heights[0][c])           // верхній ряд → Pacific
		dfs(rows-1, c, atl, heights[rows-1][c]) // нижній ряд → Atlantic
	}
	for r := 0; r < rows; r++ {
		dfs(r, 0, pac, heights[r][0])           // лівий стовпець → Pacific
		dfs(r, cols-1, atl, heights[r][cols-1]) // правий стовпець → Atlantic
	}

	// збираємо клітинки, досяжні одночасно з обох океанів (перетин множин)
	res := [][]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if pac[[2]int{r, c}] && atl[[2]int{r, c}] {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}