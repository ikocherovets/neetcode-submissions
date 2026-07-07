/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
	// якщо в нас є дві ноди, і вони мають однакові значення, йдемо далі вниз до самого дна аж поки не буде nil, це в кінці дасть нам true
    if p != nil && q != nil && p.Val == q.Val {
        return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
    }
    return false
}
