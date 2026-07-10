func climbStairs(n int) int {
    cache := make([]int, n+1)
    for i := 0; i <= n; i++ {
        cache[i] = -1
    }

    var dfs func(i int) int
    dfs = func(i int) int {
        if i >= n {
            if i == n {
                return 1
            }
            return 0
        }
        if cache[i] != -1 {
            return cache[i]
        }
        cache[i] = dfs(i + 1) + dfs(i + 2)
        return cache[i]
    }
    return dfs(0)
}