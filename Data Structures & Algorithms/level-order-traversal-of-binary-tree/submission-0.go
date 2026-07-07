/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int { // i know how to do it using bfs
	res := [][]int{}
	if root == nil {
		return res
	}
	var bfs func(root *TreeNode)

	bfs = func(root *TreeNode) {
		queue := []*TreeNode{}
		queue = append(queue, root)

		for len(queue) > 0 {
			level := []int{}
			qLen := len(queue)

			for i := 0; i < qLen; i++ {
				currentNode := queue[0]
				queue = queue[1:]
				level = append(level, currentNode.Val)

				if currentNode.Left != nil {
					queue = append(queue, currentNode.Left)
				}
				if currentNode.Right != nil {
					queue = append(queue, currentNode.Right)
				}
			}

			res = append(res, level)
		}
	}

	bfs(root)

	return res
}
