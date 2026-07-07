func isValid(s string) bool {
    stack := []rune{}
	closeToOpen := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if open, ok := closeToOpen[char]; ok {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1] 

			if top != open {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
