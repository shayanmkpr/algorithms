package main

import (
	"fmt"
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	var result []string
	var bt func(path string, node *TreeNode)
	bt = func(path string, node *TreeNode) {
		if node == nil {
			return
		}
		fmt.Println(node.Val, path)
		path += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			result = append(result, path)
		}
		path += "->"
		bt(path, node.Left)

		bt(path, node.Right)
	}
	bt("", root)
	return result
}
