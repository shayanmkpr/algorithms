package main

import "fmt"

func MinimumSeconds(nums []int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var maxGap int
	for _, target := range nums {
		fmt.Println("trying", target, nums)
		maxGap = 0
		var indices []int
		for i, element := range nums {
			if element == target {
				indices = append(indices, i)
			}
		}
		for i := 1; i < len(indices); i++ {
			maxGap = max(maxGap, indices[i]-indices[i-1])
		}
		fmt.Println(indices[0] + len(nums) - 1 - indices[len(indices)-1])
		maxGap = max(maxGap, indices[0]+len(nums)-1-indices[len(indices)-1])
	}
	return maxGap / 2
}
