/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} 
	// наступний base case - один з них не пустий через верхню умову якщо вона false
	if root == nil || subRoot == nil {
		return false
	}

	isValid := root.Val == subRoot.Val
	isLeftValid := isSameTree(root.Left, subRoot.Left)
	isRightValid := isSameTree(root.Right, subRoot.Right)

	return isValid && isLeftValid && isRightValid 
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    var dfs func(root *TreeNode, subRoot *TreeNode) bool
	dfs = func(root *TreeNode, subRoot *TreeNode) bool {
		// base case - пустий корінь
		if root == nil {
			return false
		}

		isSameTree := isSameTree(root, subRoot)
		isLeftSame := dfs(root.Left, subRoot)
		isRightSame := dfs(root.Right, subRoot)


		return isSameTree || isLeftSame || isRightSame
	}

	return dfs(root, subRoot)
}
