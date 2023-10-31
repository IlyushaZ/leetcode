package tree

// IsSymmetric is the solution to https://leetcode.com/problems/symmetric-tree/
func IsSymmetric(root *TreeNode) bool {
	return symmetric(root.Left, root.Right)
}

func symmetric(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	return node1.Val == node2.Val && symmetric(node1.Left, node2.Right) && symmetric(node1.Right, node2.Left)
}
