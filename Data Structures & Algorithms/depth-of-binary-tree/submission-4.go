/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    // base case 
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return 1 + max(leftDepth, rightDepth)
}


/* 
Пекло = ти намагаєшся симулювати весь стек у голові. Це фізично неможливо, тому боляче.
Стрибок віри = ти перевіряєш тільки дві речі (база + один крок), а решту довіряєш рекурсії, 
бо математика гарантує, що цього достатньо.
Перемикач = коли пишеш dfs(child), 
кажи "тут уже правильна відповідь, крапка" і не заглядай усередину.
*/ 
