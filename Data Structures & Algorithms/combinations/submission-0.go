func combine(n int, k int) [][]int {
    res := [][]int{}

    var backtrack func(i int, comb []int)
    backtrack = func(i int, comb []int) {
        if i > n {
            if len(comb) == k {
                res = append(res, append([]int{}, comb...))
            }
            return
        }

        comb = append(comb, i)
        backtrack(i+1, comb)
        comb = comb[:len(comb)-1]
        backtrack(i+1, comb)
    }

    backtrack(1, []int{})
    return res
}