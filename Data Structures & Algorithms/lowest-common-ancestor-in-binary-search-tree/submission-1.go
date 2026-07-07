func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	curr := root
	for curr != nil {
		if curr.Val > p.Val && curr.Val > q.Val {
			curr = curr.Left
		} else if curr.Val < p.Val && curr.Val < q.Val {
			curr = curr.Right
		} else {
			return curr 
		}
	}

	return nil
}