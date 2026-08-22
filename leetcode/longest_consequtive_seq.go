package main

/*

make a map of the existence. start from where ever, if the num + 1 existed, then continue counting, if not, count from zero.
keep the highest count.

*/

// 3 10 2 4

func longestConsecutive(nums []int) int {
	var localMax, globalMax int

	// seen := map[int]bool{}
	seen := make(map[int]bool, len(nums))

	for _, val := range nums {
		seen[val] = true
	}

	for val := range seen {
		if _, ok := seen[val-1]; !ok {
			localMax = 1
		} else {
			continue
		}
		for seen[val+1] {
			localMax++
			val++
		}
		if localMax > globalMax {
			globalMax = localMax
		}
		if globalMax == len(nums) {
			break
		}
	}

	return globalMax
}
