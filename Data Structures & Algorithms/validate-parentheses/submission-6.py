class Solution:
    def isValid(self, s: str) -> bool:
        close_to_open = {")": "(", "]": "[", "}": "{"}
        stack = []

        for char in s:
            if char in close_to_open and len(stack) > 0:
                top_bracket = stack[len(stack) - 1]
                open_from_char = close_to_open[char]

                if top_bracket != open_from_char:
                    return False

                stack = stack[:len(stack) - 1]
            else:
                stack.append(char)

        return len(stack) == 0