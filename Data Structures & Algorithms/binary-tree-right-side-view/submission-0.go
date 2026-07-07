/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    // edge case
    if root == nil {
        return []int{}
    }

    result := []int{}
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        rightSide := 0
        qLen := len(queue) // 2

        for i := 0; i < qLen; i++ {
            currentNode := queue[0]
            queue = queue[1:]

            if currentNode != nil {
                rightSide = currentNode.Val
                fmt.Println(rightSide)
                if currentNode.Left != nil {
                    queue = append(queue, currentNode.Left)
                }
                if currentNode.Right != nil {
                    queue = append(queue, currentNode.Right)
                }
            }          
        }
        result = append(result, rightSide)

    }

    return result
}
