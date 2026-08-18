package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// given a binary tree, give out all the stuff in each level, from left to right.

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		// var level []int
		level := make([]int, 0, levelSize)
		fmt.Println(level)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			level[i] = node.Val
			// level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[levelSize:]
		result = append(result, level)
	}

	return result
}
