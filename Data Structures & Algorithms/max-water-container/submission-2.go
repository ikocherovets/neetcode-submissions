func maxArea(heights []int) int {
	l, r := 0, len(heights) - 1
	maxAmount := 0

	for l < r {
		area := (r - l) * min(heights[l], heights[r])
		if maxAmount < area {
			maxAmount = area
		}
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}

	return maxAmount
}