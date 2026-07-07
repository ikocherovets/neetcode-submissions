type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var result string

	for _, str := range strs {
		encodedString := fmt.Sprintf("%d#%s", len(str), str)
		result += encodedString
	}

	return result
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}
	i := 0

	for i < len(encoded) {
		j := i

		for encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])
		word := encoded[j + 1:j + 1 + length]

		result = append(result, word)
		i = j + 1 + length
	}

	return result
}
