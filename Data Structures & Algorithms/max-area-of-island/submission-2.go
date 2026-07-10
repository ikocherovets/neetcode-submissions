func maxAreaOfIsland(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    if rows == 0 {
        return 0
    }

    var dfs func(r, c, currArea int) int
    dfs = func(r, c, currArea int) int {
        // base case
        if r < 0 || r == rows || c < 0 || c == cols || grid[r][c] == 0 {
            return 0
        }

        // що ми робимо для поточного вузла
        // рахуємо площу та топимо
        area := currArea + 1
        grid[r][c] = 0
        // leap of faith - передаю тепер площу в кожну сторону по порядку - та вибираю максимум
        area = max(area, dfs(r + 1, c, area))
        area = max(area, dfs(r - 1, c, area))
        area = max(area, dfs(r, c + 1, area))
        area = max(area, dfs(r, c - 1, area))

        return area
    }

    maxArea := 0
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 1{
                 maxArea = max(maxArea, dfs(i, j, 0))
            }
        }
    }

    return maxArea
}
