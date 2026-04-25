package main

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := nums[len(nums)/2]
	tree := &TreeNode{
		Val: root,
	}
	tree.Left = sortedArrayToBST(nums[len(nums)/2:])
	tree.Right = sortedArrayToBST(nums[:len(nums)/2])
	return tree
}
