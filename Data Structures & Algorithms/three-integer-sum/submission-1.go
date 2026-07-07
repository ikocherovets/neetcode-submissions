func threeSum(nums []int) [][]int {
	sliceCopy := append([]int{}, nums...)
    res := [][]int{}
    sort.Ints(sliceCopy)

    for i := 0; i < len(sliceCopy); i++ {
        a := sliceCopy[i]
        if a > 0 {
            break
        }
        if i > 0 && a == sliceCopy[i-1] {
            continue
        }

        l, r := i+1, len(sliceCopy)-1
        for l < r {
            threeSum := a + sliceCopy[l] + sliceCopy[r]
            if threeSum > 0 {
                r--
            } else if threeSum < 0 {
                l++
            } else {
                res = append(res, []int{a, sliceCopy[l], sliceCopy[r]})
                l++
                r--
                for l < r && sliceCopy[l] == sliceCopy[l-1] {
                    l++
                }
            }
        }
    }

    return res
}