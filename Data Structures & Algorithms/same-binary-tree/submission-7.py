# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def isSameTree(self, p, q):
        def dfs(p, q):
            if p is None and q is None:
                return True

            if p is None or q is None:
                return False

            is_node_same = p.val == q.val

            return (
                is_node_same
                and dfs(p.left, q.left)
                and dfs(p.right, q.right)
            )

        return dfs(p, q)