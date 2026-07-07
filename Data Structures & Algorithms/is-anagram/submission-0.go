func isAnagram(s string, t string) bool {
	sStore := make(map[rune]int)
	tStore := make(map[rune]int)

	for _, ch := range s {
		sStore[ch]++
	}

	for _, ch := range t {
		tStore[ch]++
	}

	if len(sStore) != len(tStore) {
		return false
	}

	for key, value := range sStore {
		if tStore[key] != value {
			return false
		}
	}

	return true
}