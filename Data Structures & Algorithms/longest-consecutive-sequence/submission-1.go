func longestConsecutive(nums []int) int {
	hashSet := make(map[int]struct{})

	for _, num := range nums {
		hashSet[num] = struct{}{} 
	}

	fmt.Println(hashSet)

	// now we have a hashSet without repeating words

	longest := 0
	for num := range hashSet {
		if _, ok := hashSet[num - 1]; ok == false { // if it's false - it is a beginning
			length := 1

			for {
				if _, ok = hashSet[num + length]; ok == true {
					length++
				} else {
					break
				}
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}
