/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		// base case 
		if root == nil {
			return 0
		}
		// якщо діаметер це сума макс ребр, то це виходить треба висота наступної ноди
		// leap of faith -> мій dfs рахує висоту поточного вузла і перевіряє чи більший діаметер за попереднє значення
		// висота вузла - це к-сть вузлів від найнижчого листка до поточного вузла
		leftHeight := dfs(root.Left)
		rightHeight := dfs(root.Right)
		
		diameter := leftHeight + rightHeight
		maxDiameter = max(maxDiameter, diameter)

		return 1 + max(leftHeight, rightHeight)
	}

	dfs(root)
	return maxDiameter
}
