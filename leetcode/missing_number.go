package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)

	seen := make(map[int]bool)
	fmt.Println(seen[0])

	for i := range nums {
		seen[nums[i]] = true
	}

	for j := 0; j <= n; j++ {
		if !seen[j] {
			return j
		}
	}

	return 0
}
