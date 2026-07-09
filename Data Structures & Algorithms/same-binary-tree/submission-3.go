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
   - перевіряє чи p.Val == q.Val
2. що вона повертає?
   - bool (true/false)

3. base case для поточного вузла
   - якщо значення обох нод nill одночасно  {
    retrun false
   }

4. що "повертають" діти? (leap of faith)
   - bool.  dfs(root.Left) та dfs(root.Right).
     Після dfs(root.Right) індукція що там вже результат
     (це і є leap of faith)

5. combine
   - мені треба глянути тепер чи p.Val == q.Val && isLeftValid && isRightValid.
*/  

func isSameTree(p *TreeNode, q *TreeNode) bool {
    var dfs func(p *TreeNode, q *TreeNode) bool 
    dfs = func(p *TreeNode, q *TreeNode) bool {
        if p == nil && q == nil {
            return true
        }
        if p != nil && q == nil || p == nil && q != nil {
            return false
        }

        isValid := p.Val == q.Val
        isLeftSame := dfs(p.Left, q.Left)
        isRightSame := dfs(p.Right, q.Right)
        return isValid && isLeftSame && isRightSame

    }

    return dfs(p, q)
}
