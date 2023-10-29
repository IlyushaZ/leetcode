package tree

// IsSameTree is the solution to https://leetcode.com/problems/same-tree
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
