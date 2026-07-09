/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
1. Контракт (що гарантує dfs для поточного вузла?)
   - dfs(root) повертає висоту дерева з коренем root.
   - Також, якщо через поточний root проходить більший діаметр,
     оновлює maxDiameter.

2. Що повертає?
   - int (висоту дерева).

3. Base case для поточного вузла
   - якщо root == nil,
     повернути 0.

4. Leap of faith (що я отримую від дітей?)
   - dfs(root.Left) повернув правильну висоту
     лівого піддерева і вже оновив maxDiameter
     всередині нього.
   - dfs(root.Right) повернув правильну висоту
     правого піддерева і вже оновив maxDiameter
     всередині нього.

5. Combine
   - знаючи висоту лівого та правого піддерев,
     рахую діаметр через поточний вузол:

        diameter = leftHeight + rightHeight

   - оновлюю maxDiameter:

        maxDiameter = max(maxDiameter, diameter)

   - повертаю висоту поточного дерева:

        1 + max(leftHeight, rightHeight)
*/ 

func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0

	var dfs func(root *TreeNode) int 
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftHeight := dfs(root.Left)
		rightHeight := dfs(root.Right)
		diameter := leftHeight + rightHeight
		maxDiameter = max(maxDiameter, diameter)

		return 1 + max(leftHeight, rightHeight)
	}

	dfs(root)

	return maxDiameter
}
