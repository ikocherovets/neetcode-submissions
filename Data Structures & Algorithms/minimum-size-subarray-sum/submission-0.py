class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        length = float('inf')
        current_sum = 0
        L = 0

        for R in range(len(nums)):
            current_sum += nums[R]

            while current_sum >= target:
                length = min(length, R - L + 1)
                current_sum -= nums[L]
                L += 1
        
        if length == float('inf'):
            return 0
        else:
            return length
        
