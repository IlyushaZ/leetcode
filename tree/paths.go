package tree

import (
	"fmt"
	"strings"
)

// BinaryTreePaths is the solution to https://leetcode.com/problems/binary-tree-paths
func BinaryTreePaths(root *TreeNode) []string {
	sb := strings.Builder{}
	ps := make([]string, 0)

	paths(root, &sb, &ps)

	return ps
}

func paths(root *TreeNode, sb *strings.Builder, ps *[]string) {
	if root == nil {
		return
	}

	sb.WriteString(fmt.Sprintf("%d", root.Val))

	// leaf is reached
	if root.Left == nil && root.Right == nil {
		*ps = append(*ps, sb.String())
		return
	}

	sb.WriteString("->")

	if root.Left != nil {
		leftSB := strings.Builder{}
		leftSB.WriteString(sb.String())

		paths(root.Left, &leftSB, ps)
	}

	if root.Right != nil {
		rightSB := strings.Builder{}
		rightSB.WriteString(sb.String())

		paths(root.Right, &rightSB, ps)
	}
}
