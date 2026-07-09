/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	/* 
	1. контракт (що робить ця функція для цього root?) - функція dfs міняє місцями дві ноди
	2. що вона повертає - нічого
	3. base case - root == nil
	4. що повертають діти - нічого (тут починається leap of faith - recursion)
	5. combine - як зліпити відповідь маючи left & right? - вже все зроблено
	*/
    if root == nil {
		return root
	}

	root.Right, root.Left = root.Left, root.Right
	invertTree(root.Right)
	invertTree(root.Left)

	return root
}
