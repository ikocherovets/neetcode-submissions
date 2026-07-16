func getConcatenation(nums []int) []int {
	ans := make([]int, len(nums) * 2)
	hashMap := make(map[int]int)
	numsLen := len(nums)
	
	for i := 0; i < numsLen; i++ {
		hashMap[i] = numsLen + i
	}

	for i1, i2 := range hashMap {
		ans[i1] = nums[i1]
		ans[i2] = nums[i1]
	}

	return ans	
}
