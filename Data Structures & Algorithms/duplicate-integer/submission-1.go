func hasDuplicate(nums []int) bool {
    numStore := make(map[int]int)

    for _, num := range nums {
        if _, ok := numStore[num]; ok {
            return true
        } else {
            numStore[num] = 1
        }
    }

    return false
}
