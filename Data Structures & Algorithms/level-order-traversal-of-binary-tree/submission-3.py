# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        result = []

        if root is None:
            return result

        def bfs(root):
            queue = deque([root])

            while len(queue) > 0:
                qLen = len(queue)
                lvlValues = []

                for i in range(qLen):
                    node = queue.popleft()
                    lvlValues.append(node.val)

                    if node.left is not None:
                        queue.append(node.left)

                    if node.right is not None:
                        queue.append(node.right)

                result.append(lvlValues)

        bfs(root)
        return result