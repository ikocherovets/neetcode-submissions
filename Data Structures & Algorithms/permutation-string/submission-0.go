func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var need, window [26]int
	for i := 0; i < len(s1); i++ {
		need[s1[i]-'a']++
		window[s2[i]-'a']++
	}
	if need == window {
		return true
	}

	for r := len(s1); r < len(s2); r++ {
		window[s2[r]-'a']++   // add new char on the right
		window[s2[r-len(s1)]-'a']-- // remove char on the left
		if need == window {
			return true
		}
	}
	return false
}