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
   - перевіряє, чи є subRoot піддеревом дерева
     з коренем root.

   - тобто:
     dfs(root) повертає true, якщо:
        або поточний root є початком subRoot
        або subRoot знаходиться зліва
        або subRoot знаходиться справа.

2. Що вона повертає?
   - bool.

3. Base case для поточного вузла
   - якщо root == nil {
        return false
     }

   - якщо root закінчився, то ми не знайшли
     місце, де починається subRoot.

4. Що "повертають" діти? (leap of faith)
   - dfs(root.Left) повертає правильну відповідь,
     чи є subRoot у лівому піддереві.

   - dfs(root.Right) повертає правильну відповідь,
     чи є subRoot у правому піддереві.

   - ми не думаємо, як вони це перевірили,
     ми довіряємо їхньому контракту.

5. Combine
   - subRoot може бути знайдений у трьох місцях:

        1. Поточний root є початком subRoot:
           isSameTree(root, subRoot)

        2. У лівому піддереві:
           dfs(root.Left)

        3. У правому піддереві:
           dfs(root.Right)

   return:
        isSameTree(root, subRoot) ||
        dfs(root.Left) ||
        dfs(root.Right)
*/


/*
1. Контракт (що робить isSameTree функція?)
   - перевіряє, чи два дерева з коренями root і subRoot
     повністю однакові.

2. Що вона повертає?
   - bool.

3. Base case для поточних вузлів
   - якщо root == nil && subRoot == nil {
        return true
     }

   - якщо один nil, а інший ні {
        return false
     }

4. Що "повертають" діти? (leap of faith)
   - isSameTree(root.Left, subRoot.Left)
     знає, чи однакові ліві піддерева.

   - isSameTree(root.Right, subRoot.Right)
     знає, чи однакові праві піддерева.

5. Combine
   - значення поточних вузлів однакові
   - ліві піддерева однакові
   - праві піддерева однакові

   return:
        root.Val == subRoot.Val &&
        isSameTree(root.Left, subRoot.Left) &&
        isSameTree(root.Right, subRoot.Right)
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var sameTree func(root *TreeNode, subRoot *TreeNode) bool 
	sameTree = func(root *TreeNode, subRoot *TreeNode) bool {
		// base case - коли воно обоє дійшли до кінця та мають однакові значення nil - вертаємо true

		if root == nil && subRoot == nil {
			return true
		}
		// наступний base case - якщо хочаби один після попередньої умови не пустий - вертаємо false
		if root == nil || subRoot == nil {
			return false
		}

		if root.Val != subRoot.Val {
			return false
		}

		isPresentLeft := sameTree(root.Left, subRoot.Left)
		isPresentRight := sameTree(root.Right, subRoot.Right)

		return isPresentLeft && isPresentRight
	}

	var dfs func(root *TreeNode, subRoot *TreeNode) bool
	dfs = func(root *TreeNode, subRoot *TreeNode) bool {
		// base case - немає дерева, де шукати subroot
		if root == nil {
			return false
		}
		if sameTree(root, subRoot) {
			return true
		}

		return dfs(root.Left, subRoot) || dfs(root.Right, subRoot)
	}

	return dfs(root, subRoot)

}
