func numIslands(grid [][]byte) int {
    directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    rows, cols := len(grid), len(grid[0])
    islands := 0

    var bfs func(r, c int)
    bfs = func(r, c int) {
        q := [][]int{{r, c}}
        grid[r][c] = '0'

        for len(q) > 0 {
            front := q[0]
            q = q[1:]
            row, col := front[0], front[1]
            for _, dir := range directions {
                nr, nc := row+dir[0], col+dir[1]
                if nr < 0 || nc < 0 || nr >= rows ||
                   nc >= cols || grid[nr][nc] == '0' {
                    continue
                }
                q = append(q, []int{nr, nc})
                grid[nr][nc] = '0'
            }
        }
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == '1' {
                bfs(r, c)
                islands++
            }
        }
    }

    return islands
}