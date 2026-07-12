/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	return 1 + max(leftHeight, rightHeight)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isBalanced(root *TreeNode) bool {
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		
		// leap of faith - результат чи збалансовані діти
		isLeftBalanced := dfs(root.Left)
		isRightBalanced := dfs(root.Right)

		leftHeight := height(root.Left)
		rightHeight := height(root.Right)

		return isLeftBalanced &&
			isRightBalanced &&
			abs(leftHeight-rightHeight) <= 1
	}

	return dfs(root)
}
