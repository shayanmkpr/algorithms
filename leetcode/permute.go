package main

import "fmt"

func permute(nums []int) [][]int {
	// given a set of integers, give out all possible permutations.
	var res [][]int
	used := make([]bool, len(nums))
	var bt func(curr []int)
	bt = func(curr []int) {
		fmt.Println(curr)
		if len(curr) == len(nums) { // means that we used every element
			// make a map of all the used paths
			temp := make([]int, len(nums))
			copy(temp, curr)
			res = append(res, temp)
			fmt.Println("done")
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			curr = append(curr, nums[i])
			used[i] = true
			bt(curr)
			// now back track to the root
			used[i] = false
			curr = curr[:len(curr)-1]
		}
	}
	bt([]int{})
	return res
}
