package tree

// IsSubtree is the solution to https://leetcode.com/problems/subtree-of-another-tree
func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if IsSameTree(root, subRoot) {
		return true
	}

	if root == nil {
		return false
	}

	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}
