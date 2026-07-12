func containsNearbyDuplicate(nums []int, k int) bool {
    mp := make(map[int]int)

    for i, num := range nums {
        if j, ok := mp[num]; ok && i-j <= k {
            return true
        }
        mp[num] = i
    }

    return false
}