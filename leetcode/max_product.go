package main

func maxProduct(nums []int) int {
	// the idea:
	// keep the max and the min at every point.
	// check if the max and the min are appropriate
	if len(nums) == 1 {
		return nums[0]
	}
	var maxHelper func(a int, b int) int
	var minHelper func(a int, b int) int

	maxHelper = func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	minHelper = func(a int, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	result := nums[0]
	maxSoFar := nums[0]
	minSoFar := nums[0]

	for i := 1; i < len(nums); i++ {

		tempMax := maxHelper(nums[i], maxHelper(maxSoFar*nums[i], minSoFar*nums[i]))
		tempMin := minHelper(nums[i], minHelper(minSoFar*nums[i], maxSoFar*nums[i]))
		maxSoFar = tempMax
		minSoFar = tempMin
		result = maxHelper(result, maxSoFar)
	}

	return result
}
