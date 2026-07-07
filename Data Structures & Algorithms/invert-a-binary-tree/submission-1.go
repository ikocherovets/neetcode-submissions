/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
		return nil
	}

	// swap the children 
	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	invertTree(root.Left) // дзеркально інвертую всі інші ноди по лівій стороні
	invertTree(root.Right) // дзеркально інвертую всі інші ноди по правій стороні

	return root
}
