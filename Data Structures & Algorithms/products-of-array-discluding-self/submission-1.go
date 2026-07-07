func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	for i, _ := range nums {
		newProduct := 1

		for j, num2 := range nums {
			if i != j {
				newProduct *= num2
			}
		}

		result[i] = newProduct
	}

	return result
}
