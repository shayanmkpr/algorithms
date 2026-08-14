package main

func canPartition(nums []int) bool {
	var half int
	for _, v := range nums {
		half += v
	}
	if half%2 == 0 {
		half = half / 2
	} else {
		return false
	}

	return false
}
