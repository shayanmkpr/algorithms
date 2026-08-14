package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(myGrid grid, start loc) {
	queue := []loc{start}
	seen := map[loc]bool{start: true}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		// some logic
		for _, neighbour := range myGrid[curr] {
			if !seen[neighbour] {
				queue = append(queue, neighbour)
				seen[neighbour] = true
				// some other logic with neigh neighbour
			}
			// some other logic with neigh neighbour
		}
	}
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var result, totalMax int
	queue := []*TreeNode{root}
	sum := map[int]int{}
	level := map[*TreeNode]int{}
	level[root] = 1
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		sum[level[curr]] += curr.Val
		for _, neighbour := range []*TreeNode{curr.Left, curr.Right} {
			if neighbour != nil && level[neighbour] == 0 {
				queue = append(queue, neighbour)
				level[neighbour] = level[curr] + 1
			}
		}
	}

	for level, levelSum := range sum {
		if levelSum > totalMax {
			totalMax = levelSum
			result = level
		}
	}
	return result
}
