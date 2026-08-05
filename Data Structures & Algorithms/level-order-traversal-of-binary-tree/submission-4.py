# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

from collections import deque

class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        result = []

        if root is None:
            return result

        def bfs(root):
            queue = deque()
            queue.append(root)

            while queue:
                level_values = []
                queue_len = len(queue)

                for _ in range(queue_len):
                    current = queue.popleft()
                    level_values.append(current.val)

                    if current.left is not None:
                        queue.append(current.left)

                    if current.right is not None:
                        queue.append(current.right)

                result.append(level_values)

        bfs(root)
        return result