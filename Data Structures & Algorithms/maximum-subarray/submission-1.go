func maxSubArray(nums []int) int {
    curSum := 0
	maxSum := nums[0]

	for _, num := range nums {
		if curSum < 0 {
			curSum = 0
		}
		curSum += num
		maxSum = max(curSum, maxSum)
	}

	return maxSum
}
