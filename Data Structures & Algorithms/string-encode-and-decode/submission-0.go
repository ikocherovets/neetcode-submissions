type Solution struct{}

const delimeter string = "#"

func (s *Solution) Encode(strs []string) string {
	var response string

	for _, str := range strs {
		encodedStr := fmt.Sprintf("%d#%s", len(str), str)
		response += encodedStr
	}

	return response 
}

func (s *Solution) Decode(encoded string) []string {
    decoded := make([]string, 0)
    i := 0

    for i < len(encoded) {
        j := i

        for string(encoded[j]) != delimeter {
            j++
        }

        length, _ := strconv.Atoi(encoded[i:j])

        word := encoded[j+1 : j+1+length]
        decoded = append(decoded, word)

        i = j + 1 + length
    }

    return decoded
}