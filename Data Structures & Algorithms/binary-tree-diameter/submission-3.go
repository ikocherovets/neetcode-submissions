/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
1. контракт (що робить dfs функція для цього root?)
   - перевіряє чи це максмальне значення діаметру (heightLeft + heigthRight)
2. що вона повертає?
   - висоту heigth ()

3. base case для поточного вузла
   - якщо root == nil {
   	return 0
   }

4. що "повертають" діти? так само висоту перевіряючи все те саме
   - int.  dfs(root.Left) та dfs(root.Right)
     (це і є leap of faith)

5. combine
   - тепер я перевіряю чи кожне піддерево має більший діаметер
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
