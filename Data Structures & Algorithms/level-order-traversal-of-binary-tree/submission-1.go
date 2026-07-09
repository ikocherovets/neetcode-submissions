/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
    
	var bfs func(root *TreeNode)
	bfs = func(root *TreeNode) {
		queue := []*TreeNode{root}

		for len(queue) > 0 {
			values := []int{}
			qLen := len(queue)

			for i := 0; i < qLen; i++ {
				node := queue[0]
				queue = queue[1:]

				values = append(values, node.Val)
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}

			res = append(res, values)
		}
	}

	bfs(root)
	return res 
}
