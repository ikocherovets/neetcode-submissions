/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    res := 0

    var dfs func(*TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }

        left := dfs(root.Left)
        right := dfs(root.Right)
        res = max(res, left + right)

        return 1 + max(left, right) // висота піддерева: беремо глибшу гілку і + 1 за поточний вузол
		// через те що це самий останній вузол, так як left / right nil, даємо +1
    }

    dfs(root)
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}