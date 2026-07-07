func productExceptSelf(nums []int) []int {
	product := make([]int, len(nums))
	fmt.Println(product)

	for i, _ := range nums {
		newProduct := 1

		for j, num2 := range nums {
			if i != j {
				newProduct *= num2
			}
		}
		fmt.Println(newProduct)

		product[i] = newProduct
	}

	return product
}
