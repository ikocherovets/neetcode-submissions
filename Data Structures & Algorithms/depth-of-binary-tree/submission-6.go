/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    /* 
	1. контракт (що робить ця функція для цього current tree?) - визначає його глибину максимальну глибину
	2. що вона повертає - число
	3. base case - root == nil -> тоді глибина 0
	4. що повертають діти - число. Максимальну глибину (тут починається leap of faith - recursion)
	5. combine - як зліпити відповідь маючи left & right? - left + right + 1
	*/
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


