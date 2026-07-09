type pair struct {
    x int
    y int
}
func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    rows, cols := len(grid), len(grid[0])
    visited := make(map[pair]bool)
    islands := 0

    var bfs func(r, c int)
    bfs = func(r,c int) {
        queue := []pair{}
        queue = append(queue, pair{r,c})

        visited[pair{r,c}] = true
        
        for len(queue) > 0 {
            p := queue[0]
            queue = queue[1:]
            dirs := [][]int{{1,0}, {-1,0}, {0,1},{0,-1}}
            for _, dir := range dirs {
                r = p.x + dir[0] // new step in row
                c = p.y + dir[1] // new step in column
                if r >= 0 && 
                    c >= 0 && 
                    r < rows && 
                    c < cols &&
                    grid[r][c] == byte('1') &&
                    !visited[pair{r,c}] {
                        visited[pair{r,c}] = true
                        queue = append(queue, pair{r,c})
                    }
            }
        }
    }

    for r := range rows {
        for c := range cols {
            if grid[r][c] == byte('1') && !visited[pair{r,c}]{
                bfs(r,c)
                islands++
            }
        }
    }
    return islands
}