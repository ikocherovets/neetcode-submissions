func getConcatenation(nums []int) []int {
    n := len(nums)
	result := make([]int, 2*len(nums))

	for i, num := range nums {
		result[i] = num
		result[i + n] = num
	}

	return result
}	
