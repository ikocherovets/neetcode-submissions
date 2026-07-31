# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:   
    def isSameTree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        if not root and not subRoot:
            return True
        
        if not root or not subRoot:
            return False
        
        isValid = root.val == subRoot.val
        isLeftValid = self.isSameTree(root.left, subRoot.left)
        isRightValid = self.isSameTree(root.right, subRoot.right)

        return isValid and isLeftValid and isRightValid

                
    def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        def dfs(root: Optional[TreeNode], subroot: Optional[TreeNode]) -> bool:
            # base case - пустий корінь
            if not root:
                return False
            
            isSameTree = self.isSameTree(root, subroot)
            isLeftSame = dfs(root.left, subroot)
            isRightSame = dfs(root.right, subroot)

            return isSameTree or isLeftSame or isRightSame

        return dfs(root, subRoot)
