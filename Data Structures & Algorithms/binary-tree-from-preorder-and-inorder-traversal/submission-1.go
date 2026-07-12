/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
1. Контракт (що робить buildTree?)
   - будує дерево з preorder + inorder.
   - повертає корінь поточного піддерева.


2. Base case:
   - якщо масиви пусті:
        return nil

   - піддерева більше немає.


3. Як знайти root?
   - preorder:
        ROOT -> LEFT -> RIGHT

   - тому:
        preorder[0] = root


4. Як знайти дітей?
   - inorder:
        LEFT -> ROOT -> RIGHT

   - знаходимо root в inorder.

   - зліва від root:
        left subtree

   - справа від root:
        right subtree


5. Leap of faith:
   - buildTree(leftPreorder, leftInorder)
     поверне правильне ліве піддерево.

   - buildTree(rightPreorder, rightInorder)
     поверне правильне праве піддерево.


6. Combine:
   - створюємо root.
   - підключаємо дітей:

        root.Left = buildTree(left)
        root.Right = buildTree(right)

   - повертаємо root.


Модель:
Я будую один вузол.
Рекурсія будує дітей.
Я з'єдную результат.
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }

    root := &TreeNode{Val: preorder[0]}

	// знаходжу preorder[0] в inorder, дістаю його індекс
    mid := 0
    for i, val := range inorder {
        if val == preorder[0] {
            mid = i
            break
        }
    }

    root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
    root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

    return root
}