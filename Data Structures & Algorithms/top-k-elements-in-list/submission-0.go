func topKFrequent(nums []int, k int) []int {
	store := make(map[int]int)
	type Value struct {
		number 	int
		count 	int
	}

	for _, num := range nums {
		store[num]++
	}

	fmt.Println(store)

	var structSlice []Value
	fmt.Println("slice", structSlice)
	for k, v := range store {
		structSlice = append(structSlice, Value{k, v})
	}

	fmt.Printf("slice: %+v", structSlice)

	sort.Slice(structSlice, func(i, j int) bool {
		return structSlice[i].count > structSlice[j].count
	})

	mostFrequentElements := structSlice[:k]
	var result []int

	for _, element := range mostFrequentElements {
		result = append(result, element.number)
	}

	return result

}
