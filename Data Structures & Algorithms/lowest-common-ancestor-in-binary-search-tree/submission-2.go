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

/*If both values are smaller than the current node -> both must lie in the left subtree.
If both values are greater than the current node -> both must lie in the right subtree.
Otherwise, the current node is the split point where one node is on the left and the other is on the right (or one is equal to the current node).
That split point is the Lowest Common Ancestor (LCA).
*/