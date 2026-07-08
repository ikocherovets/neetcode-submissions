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
========================
RECURSION CHECKLIST
========================

❌ Пекло
Не симулюй увесь стек викликів у голові.
Це фізично неможливо й лише заплутує.

✅ Jump of Faith (Стрибок віри)
Припусти, що dfs(child) уже повернув правильну відповідь.
Не думай, як саме він її отримав.

Твоя відповідальність — лише поточний вузол.
Ти використовуєш відповіді лише від одного рівня нижче.

1. Base Case
   Коли рекурсія повинна зупинитися?

2. Jump of Faith
   dfs(node.Left)  -> вже правильна відповідь.
   dfs(node.Right) -> вже правильна відповідь.

3. Combine
   Як із цих відповідей отримати відповідь
   для поточного вузла?

4. Return
   Поверни тільки те, що обіцяє контракт функції.
   Не пояснюй, як рекурсія дійшла до цих значень.

Мантра:
"Trust the recursive call. Solve only the current problem."
("Довірся рекурсивному виклику. Розв'язуй лише поточну задачу.")
*/
