func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	
	clean := ""
	for _, ch := range strings.ToLower(s) {
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
    		clean += string(ch)
		}
	}

	left, right := 0, len(clean)-1
	for left < right {
		if clean[left] != clean[right] {
			return false
		}
		left++
		right--
	}
	return true
}