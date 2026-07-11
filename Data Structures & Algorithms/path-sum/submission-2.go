/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	var dfs func(root *TreeNode, curSum int) bool
	dfs = func(root *TreeNode, curSum int) bool {
		// base case - якщо передаю пусте дерево
		if root == nil {
			return false
		}
		if root.Left == nil && root.Right == nil {
			if curSum + root.Val == targetSum {
				return true
			}
		}

		return dfs(root.Left, curSum + root.Val) || dfs(root.Right, curSum + root.Val)
	}

	return dfs(root, 0)
}
