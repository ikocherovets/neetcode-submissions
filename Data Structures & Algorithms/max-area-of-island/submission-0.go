func maxAreaOfIsland(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    visited := make([][]bool, rows)
    for i := range visited {
        visited[i] = make([]bool, cols)
    }

    bfs := func(sr, sc int) int {
        q := [][2]int{{sr, sc}}
        visited[sr][sc] = true
        area := 0

        for len(q) > 0 {
            cell := q[len(q)-1]
            q = q[:len(q)-1]
            r, c := cell[0], cell[1]
            area++

            for _, d := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
                nr, nc := r+d[0], c+d[1]
                if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 && !visited[nr][nc] {
                    visited[nr][nc] = true
                    q = append(q, [2]int{nr, nc})
                }
            }
        }
        return area
    }

    area := 0
    for r := range grid {
        for c := range grid[r] {
            if grid[r][c] == 1 && !visited[r][c] {
                area = max(area, bfs(r, c))
            }
        }
    }
    return area
}