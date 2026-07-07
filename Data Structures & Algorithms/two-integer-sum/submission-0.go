func twoSum(nums []int, target int) []int {
    store := make(map[int]int) // number -> index
	
	for i, number := range nums {
		complement := target - number

		if j, ok := store[complement]; ok == true {
			return []int{j, i}
		}

		store[number] = i
	}

	return []int{0, 0}
}
