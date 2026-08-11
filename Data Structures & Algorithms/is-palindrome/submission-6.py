class Solution:
    def isPalindrome(self, s: str) -> bool:
        norm_str = ''
        for char in s:
            if char.lower() >= 'a' and char.lower() <= 'z':
                norm_str += char.lower()
            if char >= '0' and char <= '9':
                norm_str += char
        
        L, R = 0, len(norm_str) - 1
        while L < R:
            if norm_str[L] != norm_str[R]:
                return False
            L += 1
            R -= 1
        
        return True