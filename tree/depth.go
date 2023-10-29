package tree

// MinDepthRecursive is the solution to https://leetcode.com/problems/minimum-depth-of-binary-tree
// This solution uses recursive approach (DFS)
func MinDepthRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := MinDepthRecursive(root.Left) + 1
	rd := MinDepthRecursive(root.Right) + 1

	if ld == 1 {
		return rd
	} else if rd == 1 {
		return ld
	}

	return min(ld, rd)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MinDepthBFS is the solution to https://leetcode.com/problems/minimum-depth-of-binary-tree
// This solution uses BFS approach with queue
func MinDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	queueLen := 1
	depth := 0

	for queueLen > 0 {
		depth++

		for i := 0; i < queueLen; i++ {
			if queue[i].Left == nil && queue[i].Right == nil {
				return depth
			}

			if l := queue[i].Left; l != nil {
				queue = append(queue, l)
			}

			if r := queue[i].Right; r != nil {
				queue = append(queue, r)
			}
		}

		queue = queue[queueLen:]
		queueLen = len(queue)
	}

	return depth
}

// MaxDepth is the solution to https://leetcode.com/problems/maximum-depth-of-binary-tree
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ld := MaxDepth(root.Left) + 1
	rd := MaxDepth(root.Right) + 1

	if ld > rd {
		return ld
	}

	return rd
}
