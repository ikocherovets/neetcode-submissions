/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    var dfs func(p *TreeNode, q* TreeNode) bool
	dfs = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if p != nil && q == nil || p == nil && q != nil {
            return false
        }

		isValid := p.Val == q.Val

		return isValid && dfs(p.Left, q.Left) && dfs(p.Right, q.Right)
	}

	return dfs(p, q)
}
