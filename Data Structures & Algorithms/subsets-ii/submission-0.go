func subsetsWithDup(nums []int) [][]int {
    var res [][]int
    sort.Ints(nums)

    var backtrack func(int, []int)
    backtrack = func(i int, subset []int) {
        if i == len(nums) {
            res = append(res, append([]int{}, subset...))
            return
        }

        subset = append(subset, nums[i])
        backtrack(i+1, subset)
        subset = subset[:len(subset)-1]

        for i+1 < len(nums) && nums[i] == nums[i+1] {
            i++
        }

        backtrack(i+1, subset)
    }

    backtrack(0, []int{})
    return res
}