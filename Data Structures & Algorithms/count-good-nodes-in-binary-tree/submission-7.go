/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    if root == nil {
		return 0
	}

	counter := 0

	var dfs func(root *TreeNode, curMaxVal int)
	dfs = func(root *TreeNode, curMaxVal int) {
		if root == nil {
			return 
		}

		if root.Val >= curMaxVal {
			counter++
		}
		curMaxVal = max(curMaxVal, root.Val)

		dfs(root.Left, curMaxVal)
		dfs(root.Right, curMaxVal)
	}

	dfs(root, root.Val)

	return counter
}
