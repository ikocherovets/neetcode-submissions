/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 /*
0. Чи потрібна поточному вузлу інформація від батьків?

   ні, наразі не бачу

------------------------------------------------------------

1. Контракт (що робить dfs для поточного вузла?)

   - пушить результат в масив res

------------------------------------------------------------

2. Що повертає?

   нічого

------------------------------------------------------------

3. Base case для поточного вузла

   if root == nil {
       return
   }

------------------------------------------------------------

4. Leap of faith (що роблять діти?)
	запускаю все так само

------------------------------------------------------------

5. Combine (що роблю для поточного вузла?)
	нічого, це inorder dfs


*/
func kthSmallest(root *TreeNode, k int) int {
	result := []int{}
    var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		// base case 
		if root == nil {
       		return
   		}
		dfs(root.Left)
		result = append(result, root.Val)
		dfs(root.Right)
	}

	dfs(root)
	return result[k - 1]
}
