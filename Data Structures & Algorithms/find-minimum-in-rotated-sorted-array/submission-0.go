func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1 // мінімум праворуч від mid
		} else {
			right = mid // мінімум тут або ліворуч
		}
	}

	return nums[left] // left == right — це і є мінімум
}