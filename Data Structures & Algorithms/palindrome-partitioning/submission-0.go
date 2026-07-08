func partition(s string) [][]string {
    res := [][]string{}
    part := []string{}

    var dfs func(j, i int)
    dfs = func(j, i int) {
        if i >= len(s) {
            if i == j {
                res = append(res, append([]string{}, part...))
            }
            return
        }

        if isPali(s, j, i) {
            part = append(part, s[j:i+1])
            dfs(i+1, i+1)
            part = part[:len(part)-1]
        }

        dfs(j, i+1)
    }

    dfs(0, 0)
    return res
}

func isPali(s string, l, r int) bool {
    for l < r {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    return true
}