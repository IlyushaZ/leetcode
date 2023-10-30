package tree

// IsBalanced is the solution to https://leetcode.com/problems/balanced-binary-tree
func IsBalanced(root *TreeNode) bool {
	_, balanced := isBalanced(root)
	return balanced
}

// isBalanced counts maxDepth of binary tree and checks whether it's balanced on every level.
// In case when unbalanced subtree is detected -1 is returned as depth.
func isBalanced(root *TreeNode) (depth int, balanced bool) {
	if root == nil {
		return 0, true
	}

	lDepth, lBalanced := isBalanced(root.Left)
	if !lBalanced {
		return -1, false
	}

	rDepth, rBalanced := isBalanced(root.Right)
	if !rBalanced {
		return -1, false
	}

	maxDepth := max(lDepth, rDepth) + 1

	if abs(lDepth-rDepth) > 1 {
		return -1, false
	}

	return maxDepth, true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
