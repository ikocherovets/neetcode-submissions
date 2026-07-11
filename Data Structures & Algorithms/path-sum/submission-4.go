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

		curSum += root.Val
		if root.Left == nil && root.Right == nil {
			return curSum == targetSum
		}

		return dfs(root.Left, curSum) || dfs(root.Right, curSum)
	}

	return dfs(root, 0)
}
