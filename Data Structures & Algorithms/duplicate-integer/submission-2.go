func hasDuplicate(nums []int) bool {
    store := make(map[int]int) // []int{1,2,3,4,5,5}

    for _, num := range nums {
        if _, ok := store[num]; ok == true {
            return true
        } else {
            store[num]++
        }
    }

    return false
}
