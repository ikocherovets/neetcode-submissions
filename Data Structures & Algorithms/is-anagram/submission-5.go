func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    type charStore map[byte]int

    sCharStore := make(charStore)
    tCharStore := make(charStore)

    for i, char := range s {
        sCharStore[byte(char)]++
		tCharStore[byte(t[i])]++
    }

	for k, sCharValue := range sCharStore {
		tCharValue := tCharStore[k]
		if sCharValue != tCharValue {
			return false
		}
	}

	return true
}
