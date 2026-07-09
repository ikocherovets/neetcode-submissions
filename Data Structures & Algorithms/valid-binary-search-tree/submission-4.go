/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
 1. контракт - що робить функція для поточного вузла? - перевіряє чи він знаходиться в заданих межах (min < val < max)
 2. що повертає? - bool 
 3. base case - якщо ми в самому низу - значить все гуд і вертаємо true
 3. діти - запускають так само в себе перевірку
 4. combine -  
 */

func isValidBST(root *TreeNode) bool {
	min, max := -1001, 1001
	var dfs func(root *TreeNode, min, max int) bool
	dfs = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}

		fmt.Printf("min is %d\n", min)
		fmt.Printf("max is %d\n", max)

		if root.Val >= max || root.Val <= min {
			return false
		}

		isLeftOK := dfs(root.Left, min, root.Val)
		isRightOK := dfs(root.Right, root.Val, max)

		return isLeftOK && isRightOK
	}

    return dfs(root, min, max)

}
