func topKFrequent(nums []int, k int) []int {
	store := make(map[int]int)
	type Value struct {
		number 	int
		count 	int
	}

	for _, num := range nums {
		store[num]++
	}

	structSlice := make([]Value, 0, len(store)) // 0 capacity means that we have [] slice, it is empty
	// if you write make(interface{}, len(store)), it will have default values [any, any]
	fmt.Println("slice", structSlice)
	for k, v := range store {
		structSlice = append(structSlice, Value{k, v})
	}

	sort.Slice(structSlice, func(i, j int) bool {
		return structSlice[i].count > structSlice[j].count
	})

	
	result := make([]int, k)

	for i := 0; i < k; i++ {
		result[i] = structSlice[i].number
	}

	return result

}
