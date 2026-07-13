func islandsAndTreasure(grid [][]int) {
	rows, cols := len(grid), len(grid[0])
	visit := make(map[[2]int]bool)
	queue := [][2]int{} // було [][2]int{{}} — це додавало зайвий {0,0}

	addRoom := func(r, c int) {
		if r < 0 || c < 0 || r == rows || c == cols || visit[[2]int{r, c}] || grid[r][c] == -1 {
			return
		}
		visit[[2]int{r, c}] = true
		queue = append(queue, [2]int{r, c})
	}

	// 1. знаходимо всі скарби (джерела для multi-source BFS)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 {
				queue = append(queue, [2]int{r, c})
				visit[[2]int{r, c}] = true
			}
		}
	}

	// 2. BFS хвилями, рахуємо відстань
	dist := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			r, c := cur[0], cur[1]

			grid[r][c] = dist

			addRoom(r+1, c)
			addRoom(r-1, c)
			addRoom(r, c+1)
			addRoom(r, c-1)
		}
		dist++
	}
}