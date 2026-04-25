package main

import "fmt"

func maxDepth(root *TreeNode) int {
	// classical DFS
	maxDepth := 0
	depth := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			fmt.Println("we got to the bottom")
			depth = 0
			if depth > maxDepth {
				maxDepth = depth
			}
			return
		} else {
			depth++
		}

		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return maxDepth
}
