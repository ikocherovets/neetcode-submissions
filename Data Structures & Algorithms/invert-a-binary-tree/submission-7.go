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
   - міняє місцями left and right

2. що вона повертає?
   - нічого (змінює дерево in-place)

3. base case
   - root == nil -> return (кінець інвертування)

4. що "повертають" діти? (leap of faith)
   - інвертують в себе все вузли

5. combine
   - нічого складати не треба. Просто міняємо місцями дітей
     (або до рекурсії, або після — обидва варіанти правильні).
*/ 

func invertTree(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		root.Left, root.Right = root.Right, root.Left
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	return root
}
