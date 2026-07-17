func majorityElement(nums []int) int {
	frequencyMap := make(map[int]int)

	for _, num := range nums {
		frequencyMap[num]++
	}

	for k, v := range frequencyMap {
		if v > len(nums) / 2 {
			return k
		}
	}
	
	return 0
}
