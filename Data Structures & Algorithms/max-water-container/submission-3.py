class Solution:
    def maxArea(self, heights: List[int]) -> int:
        l, r = 0, len(heights) - 1
        max_area = 0

        while l < r:
            area = min(heights[r], heights[l]) * (r - l)
            if area > max_area:
                max_area = area
            if heights[r] > heights[l]:
                l += 1
            else:
                r -= 1

        return max_area