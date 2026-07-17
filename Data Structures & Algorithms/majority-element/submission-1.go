func majorityElement(nums []int) int {
    count := make(map[int]int)
    res, maxCount := 0, 0

    for _, num := range nums {
        count[num]++
        if count[num] > maxCount {
            res = num
            maxCount = count[num]
        }
    }
    return res
}