/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    result := []int{}
	if root == nil {
		return result
	}

	var bfs func(root *TreeNode) 
	bfs = func(root *TreeNode) {
		queue := []*TreeNode{root}

		for len(queue) > 0 {
			qLen := len(queue)
			rightSide := 0

			for i := 0; i < qLen; i++ {
				node := queue[0]
				queue = queue[1:]

				rightSide = node.Val
				
				if node.Left != nil {
					queue = append(queue, node.Left)
				}

				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}

			result = append(result, rightSide)
		}
	}

	bfs(root)	

	return result
}
