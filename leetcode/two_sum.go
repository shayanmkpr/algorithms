package main

/*
we have a set of numberss. give the indeices of those that sum up to target.
*/

func twoSum(numbers []int, target int) []int {
	var l, h int = 0, len(numbers) - 1
	for l < h {
		val := numbers[l] + numbers[h]
		if val == target {
			return []int{l + 1, h + 1}
		}
		if val > target {
			h--
		}
		if val < target {
			l++
		}
	}
	return nil
}
