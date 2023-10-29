package tree

// MergeTrees is the solution to https://leetcode.com/problems/merge-two-binary-trees
func MergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}

	return &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  MergeTrees(root1.Left, root2.Left),
		Right: MergeTrees(root1.Right, root2.Right),
	}
}
