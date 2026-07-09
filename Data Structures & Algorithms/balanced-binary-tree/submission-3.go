/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
 1. Контракт
   dfs(root) повертає, чи дерево з коренем root balanced.

2. Що повертає?
   bool.

3. Base case
   root == nil -> true.

4. Leap of faith
   dfs(root.Left) повертає, чи ліве піддерево balanced.
   dfs(root.Right) повертає, чи праве піддерево balanced.

5. Combine
   Поточне дерево balanced, якщо:
   - ліве balanced
   - праве balanced
   - abs(leftHeight - rightHeight) <= 1
 */

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	return 1 + max(leftHeight, rightHeight)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isBalanced(root *TreeNode) bool {
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true
		}

		// Leap of faith:
		// діти вже знають, чи їхні піддерева balanced
		isLeftBalanced := dfs(root.Left)
		isRightBalanced := dfs(root.Right)

		// Combine:
		// перевіряємо висоту поточного вузла
		leftHeight := height(root.Left)
		rightHeight := height(root.Right)

		return isLeftBalanced &&
			isRightBalanced &&
			abs(leftHeight-rightHeight) <= 1
	}

	return dfs(root)
}