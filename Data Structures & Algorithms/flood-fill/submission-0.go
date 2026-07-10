func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    orig := image[sr][sc]
    if orig == color {
        return image
    }
    m, n := len(image), len(image[0])

    var dfs func(r, c int)
    dfs = func(r, c int) {
        if r < 0 || r >= m || c < 0 || c >= n || image[r][c] != orig {
            return
        }
        image[r][c] = color
        dfs(r+1, c)
        dfs(r-1, c)
        dfs(r, c+1)
        dfs(r, c-1)
    }

    dfs(sr, sc)
    return image
}