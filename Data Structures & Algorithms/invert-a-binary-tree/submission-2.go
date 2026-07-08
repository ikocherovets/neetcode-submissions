/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	// base case щоб зупинити рекурсію
	if root == nil {
		return nil
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	
	invertTree(root.Left)
	invertTree(root.Right)

	return root
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


ТВОЯ ВІДПОВІДАЛЬНІСТЬ - ЛИШЕ ПОТОЧНИЙ ВУЗОЛ!!!
ТИ ВИКОРИСТОВУЄШ ВІДПОВІДІ ЛИШЕ ВІД ОДНОГО РІВНЯ НИЖЧЕ.

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

