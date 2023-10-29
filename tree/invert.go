package tree

// InvertTree is the solution to https://leetcode.com/problems/invert-binary-tree
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	return &TreeNode{
		Val:   root.Val,
		Left:  InvertTree(root.Right),
		Right: InvertTree(root.Left),
	}
}
