func maxAreaOfIsland(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])

    var dfs func(r, c int) int
    dfs = func(r, c int) int {
        if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == 0 {
            return 0
        }
        grid[r][c] = 0 // mark visited
        return 1 + dfs(r+1, c) + dfs(r-1, c) + dfs(r, c+1) + dfs(r, c-1)
    }

    maxArea := 0
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 1 {
                maxArea = max(maxArea, dfs(r, c))
            }
        }
    }
    return maxArea
}