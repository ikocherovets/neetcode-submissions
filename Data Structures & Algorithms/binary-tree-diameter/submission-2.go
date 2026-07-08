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

        // Діаметр, що проходить через поточний вузол
        res = max(res, left+right)

        // Повертаємо МАКСИМАЛЬНУ висоту піддерева, треба для того щоб рахувати батьківська діаметри
        return 1 + max(left, right)
    }

    dfs(root)
    return res
}


