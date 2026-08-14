package main

func search(nums []int, target int) int {
	// find 0
	// 4 5 6 7 0 1 2 3
	l, h := 0, len(nums)-1
	for l <= h {
		curr := l + (h-l)/2
		if nums[curr] == target {
			return curr
		}
		if nums[l] <= nums[curr] { // curr is in the left half
			if nums[curr] > target && nums[l] <= target {
				h = curr - 1
			} else {
				l = curr + 1
			}
		} else {
			if nums[curr] < target && nums[h] >= target {
				l = curr + 1
			} else {
				h = curr - 1
			}
		}
	}
	return -1
}
