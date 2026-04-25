package main

func sortColors(nums []int) []int {
	// it should be in place

	low, mid, high := 0, 0, len(nums)-1 // three pointers
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[mid], nums[low] = nums[low], nums[mid]
			mid++
			low++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}

	return nums
}
