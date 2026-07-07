/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    var dfs func(node *TreeNode, curSum int) bool
    dfs = func(node *TreeNode, curSum int) bool {
        if node == nil {
            return false
        }

        curSum += node.Val
        if node.Left == nil && node.Right == nil {
            return curSum == targetSum
        }

        return dfs(node.Left, curSum) || dfs(node.Right, curSum)
    }

    return dfs(root, 0)
}