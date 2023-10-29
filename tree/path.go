package tree

// HasPathSum is the solution to https://leetcode.com/problems/path-sum
func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}

	return HasPathSum(root.Left, targetSum-root.Val) || HasPathSum(root.Right, targetSum-root.Val)
}
