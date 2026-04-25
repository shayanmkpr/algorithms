package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	sort.Ints(nums)

	var diff int
	var sum int
	var tempSum int
	var tempDiff int

	// setting initial values to non Zero values
	sum = nums[0] + nums[1] + nums[2]
	if sum > target {
		diff = sum - target
	} else {
		diff = target - sum
	}

	for i := 0; i < len(nums); i++ {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			tempSum = nums[i] + nums[l] + nums[r]
			if tempSum > target {
				tempDiff = tempSum - target
				r--
				if nums[r] == nums[r-1] {
					r--
				}
			} else {
				tempDiff = target - tempSum
				l++
				if nums[l] == nums[l-1] {
					l++
				}
			}
			if tempDiff == 0 {
				return tempSum
			}
			if tempDiff < diff {
				diff = tempDiff
				sum = tempSum
			}
		}
	}
	return sum
}
