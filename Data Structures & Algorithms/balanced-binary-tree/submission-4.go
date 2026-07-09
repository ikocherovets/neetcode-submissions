/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
1. Контракт (що робить dfs функція для поточного вузла?)
   - повертає, чи є дерево з коренем root balanced.
   - balanced означає:
        abs(leftHeight - rightHeight) <= 1
     і всі піддерева також balanced.

2. Що вона повертає?
   - bool.

3. Base case для поточного вузла
   - якщо root == nil {
        return true
     }
   - порожнє дерево завжди balanced.

4. Що "повертають" діти? (leap of faith)
   - dfs(root.Left) повертає правильну відповідь,
     чи ліве піддерево balanced.
   - dfs(root.Right) повертає правильну відповідь,
     чи праве піддерево balanced.

   - ми не думаємо, як вони це перевірили,
     ми довіряємо їхньому контракту.

5. Combine
   - щоб поточне дерево було balanced:
        1. ліве піддерево balanced
        2. праве піддерево balanced
        3. різниця висот <= 1

   return:
        isLeftBalanced &&
        isRightBalanced &&
        abs(leftHeight - rightHeight) <= 1
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