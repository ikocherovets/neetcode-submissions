class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        closeToOpen = {
            ")": "(",
            "]": "[",
            "}": "{",
        }

        for c in s:
            if c in closeToOpen:
                if len(stack) > 0 and stack[len(stack)-1] == closeToOpen[c]:
                    stack = stack[:len(stack)-1]
                else:
                    return False
            else:
                stack.append(c)

        return len(stack) == 0