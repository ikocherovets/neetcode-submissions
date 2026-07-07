func isValid(s string) bool {
    stack := []rune{}
    closeToOpen := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for _, c := range s {
        if open, ok := closeToOpen[c]; ok {
            // No matching opening bracket
            if len(stack) == 0 {
                return false
            }

            // Pop
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            if top != open {
                return false
            }
        } else {
            // Push
            stack = append(stack, c)
        }
    }

    return len(stack) == 0
}