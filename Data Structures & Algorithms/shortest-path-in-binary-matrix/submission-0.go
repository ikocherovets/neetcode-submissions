func shortestPathBinaryMatrix(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return -1
	}
	cols := len(grid[0])

	// Start or end is blocked.
	if grid[0][0] == 1 || grid[rows-1][cols-1] == 1 {
		return -1
	}

	dirs := [8][2]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	queue := [][2]int{{0, 0}}
	visited := make(map[[2]int]bool)
	visited[[2]int{0, 0}] = true

	length := 1

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]

			r, c := cur[0], cur[1]

			if r == rows-1 && c == cols-1 {
				return length
			}

			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				next := [2]int{nr, nc}

				if nr >= 0 && nr < rows &&
					nc >= 0 && nc < cols &&
					grid[nr][nc] == 0 &&
					!visited[next] {

					visited[next] = true
					queue = append(queue, next)
				}
			}
		}

		length++
	}

	return -1
}