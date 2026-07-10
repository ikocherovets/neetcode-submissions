func maxAreaOfIsland(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    if rows == 0 {
        return 0
    }

    var dfs func(r, c int) int
    dfs = func(r, c int) int {
        // base case
        if r < 0 || r == rows || c < 0 || c == cols || grid[r][c] == 0 {
            return 0
        }

        // що ми робимо для поточного вузла
        // рахуємо площу та топимо
        grid[r][c] = 0
        // leap of faith - сусідні клітинки повертають площу яка прилягає до поточної клітинки яка є 1
        return 1 + dfs(r + 1, c) + dfs(r - 1, c) + dfs(r, c + 1) + dfs(r, c - 1)
    }

    maxArea := 0
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 1 {
                maxArea = max(maxArea, dfs(i, j))
            }
        }
    }

    return maxArea
}
