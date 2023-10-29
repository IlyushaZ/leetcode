package tree

// AverageOfLevels is the solution to https://leetcode.com/problems/average-of-levels-in-binary-tree
func AverageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)

	queue := []*TreeNode{root}
	queueLen := 1

	for queueLen > 0 {
		currSum := sum(queue[:queueLen]...)
		res = append(res, float64(currSum)/float64(queueLen))

		// add children to queue
		for i := 0; i < queueLen; i++ {
			if l := queue[i].Left; l != nil {
				queue = append(queue, l)
			}

			if r := queue[i].Right; r != nil {
				queue = append(queue, r)
			}
		}

		// remove current level's nodes from queue
		queue = queue[queueLen:]
		queueLen = len(queue)
	}

	return res
}

func sum(nodes ...*TreeNode) (s int) {
	for _, n := range nodes {
		s += n.Val
	}
	return
}
