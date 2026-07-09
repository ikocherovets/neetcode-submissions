/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
/*
1. Контракт (що робить dfs для поточного вузла?)
   - перевіряє, чи поточний root є good node,
     знаючи максимальне значення на шляху
     від початкового root до цього вузла.

   - після цього передає оновлений максимум
     своїм дітям.

2. Що повертає?
   - нічого.
   - відповідь накопичується у counter.

3. Base case для поточного вузла
   - якщо root == nil {
        return
     }

4. Leap of faith (що роблять діти?)
   - dfs(root.Left, currentMax)
     правильно порахує good nodes
     у лівому піддереві.

   - dfs(root.Right, currentMax)
     правильно порахує good nodes
     у правому піддереві.

5. Combine
   - якщо root.Val >= currentMax:
        counter++

   - оновлюємо максимум:

        currentMax = max(currentMax, root.Val)

   - передаємо цей максимум дітям.
*/

func goodNodes(root *TreeNode) int {
    counter := 0
	var dfs func(root *TreeNode, curMaxVal int)
	dfs = func(root *TreeNode, curMaxVal int) {
		// base case, якщо ми в самомому дні
		if root == nil {
			return 
		}

		// що ми робимо для поточного вузла? - перевіряємо чи він більший за попередній батьківський максимум
		if root.Val >= curMaxVal {
			counter++
		}
		curMaxVal = max(curMaxVal, root.Val)

		// далі запускаємо для лівого дерева та правого те саме
		dfs(root.Left, curMaxVal)
		dfs(root.Right, curMaxVal)
	}

	dfs(root, root.Val)

	return counter
}
