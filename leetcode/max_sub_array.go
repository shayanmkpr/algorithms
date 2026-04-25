package main

import "fmt"

func maxSubArray(nums []int) int {
	fmt.Println(nums)
	tempMax := nums[0]
	maxSoFar := nums[0]

	for i := 1; i < len(nums); i++ {
		if tempMax < 0 {
			tempMax = nums[i]
		} else {
			tempMax += nums[i]
		}
		if tempMax > maxSoFar {
			maxSoFar = tempMax
		}
	}
	return maxSoFar
}
