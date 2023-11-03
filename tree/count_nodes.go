package tree

// CountNodes is the solution to https://leetcode.com/problems/count-complete-tree-nodes
func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	queueLen := 1
	cnt := 1

	for queueLen > 0 {
		for i := 0; i < queueLen; i++ {
			if l := queue[i].Left; l != nil {
				queue = append(queue, l)
			}

			if r := queue[i].Right; r != nil {
				queue = append(queue, r)
			}
		}

		queue = queue[queueLen:]
		queueLen = len(queue)
		cnt += queueLen
	}

	return cnt
}
