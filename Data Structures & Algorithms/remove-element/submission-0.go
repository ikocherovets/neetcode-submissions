func removeElement(nums []int, val int) int {
    temp := make([]int, 0)

	for _, num := range nums {
		if num != val {
			temp = append(temp, num)
		}		
	}

	for i := 0; i < len(temp); i++ {
		nums[i] = temp[i]
	}

	return len(temp)
}
