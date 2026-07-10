func islandsAndTreasure(grid [][]int) {
    rows, cols := len(grid), len(grid[0])
    directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    INF := 2147483647

    bfs := func(r, c int) int {
        q := [][2]int{{r, c}}
        visit := make([][]bool, rows)
        for i := range visit {
            visit[i] = make([]bool, cols)
        }
        visit[r][c] = true
        steps := 0

        for len(q) > 0 {
            size := len(q)
            for i := 0; i < size; i++ {
                current := q[0]
                q = q[1:]
                row, col := current[0], current[1]
                if grid[row][col] == 0 {
                    return steps
                }
                for _, dir := range directions {
                    nr, nc := row+dir[0], col+dir[1]
                    if nr >= 0 && nc >= 0 && nr < rows && nc < cols &&
                       !visit[nr][nc] && grid[nr][nc] != -1 {
                        visit[nr][nc] = true
                        q = append(q, [2]int{nr, nc})
                    }
                }
            }
            steps++
        }
        return INF
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == INF {
                grid[r][c] = bfs(r, c)
            }
        }
    }
}