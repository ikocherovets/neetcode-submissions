/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftDepth := dfs(root.Left)
		rightDepth := dfs(root.Right)

		return 1 + max(leftDepth, rightDepth)
	}

	return dfs(root)
}
