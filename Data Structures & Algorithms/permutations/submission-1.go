func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	perms := permute(nums[1:])
	result := make([][]int, 0)

	for _, perm := range perms {
		for pos := 0; pos <= len(perm); pos++ {
			result = append(result, insert(perm, nums[0], pos))
		}
	}

	return result
}

func insert(arr []int, value int, pos int) []int {
	result := make([]int, 0, len(arr)+1)

	result = append(result, arr[:pos]...)
	result = append(result, value)
	result = append(result, arr[pos:]...)

	return result
}