package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	used := map[int]bool{}
	var path []int
	var bt func()
	bt = func() {
		if len(path) == len(nums) {
			// append path
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
		}
		for i := range nums {
			// fmt.Println(i)
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			bt()
			// roll back a single step
			fmt.Println(len(path), path, nums[i], i)
			if len(path) > 0 {

				path = path[:len(path)-1]
				used[i] = false
			}
		}
	}
	bt()
	return result
}
