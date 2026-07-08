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
		// якщо старт нашого рута пустий
		if root == nil {
			return false
		}

		curSum += root.Val

		// якщо ми прийшли в самий кінець - перевіряємо
		if root.Left == nil && root.Right == nil {
			if curSum == targetSum {
				return true
			} else {
				return false
			}
		}

		return (dfs(root.Left, curSum) || dfs(root.Right, curSum))

	}

	return dfs(root, 0)
}
