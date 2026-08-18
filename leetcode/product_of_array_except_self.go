package main

func productExceptSelf(nums []int) []int {
	if len(nums) <= 0 {
		return []int{}
	}
	result := make([]int, len(nums))
	n := len(nums) - 1
	l, r := make([]int, len(nums)), make([]int, len(nums))
	l[0], r[n] = 1, 1
	for i := 1; i < len(nums); i++ {
		l[i] = nums[i-1] * l[i-1]
		r[n-i] = nums[n-i+1] * r[n-i+1]
	}
	for i := 0; i < len(nums); i++ {
		result[i] = l[i] * r[i]
	}
	return result
}
