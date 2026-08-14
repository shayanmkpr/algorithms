package main

func minimizeArrayValue(nums []int) int {
	/*
		find the max value, and the minimum behind it, ignore the rest.
		the answer is the mean of those.
	*/
	var maxVal, maxIdx, minVal, minIdx int
	for minVal <= maxVal {

		for i, val := range nums {
			if val >= maxVal {
				maxVal = val
				maxIdx = i
			}
		}
		nums = nums[:maxIdx+1]
		minVal = maxVal
		minIdx = maxIdx
		for i, val := range nums {
			if val < minVal {
				minVal = val
				minIdx = i
			}
		}
		nums[maxIdx]++
		nums[minIdx]--
	}
	return maxVal
}
