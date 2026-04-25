package main

func quickSelect(arr []int, k int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		pivotIndex := left
		pivotValue := arr[right] // Use last element as pivot
		for i := left; i < right; i++ {
			if arr[i] <= pivotValue {
				arr[i], arr[pivotIndex] = arr[pivotIndex], arr[i]
				pivotIndex++
			}
		}
		arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]
		if pivotIndex == k {
			return arr[k]
		} else if k < pivotIndex {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}

	return arr[k] // Should never reach here with valid input
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	if len(nums)%2 == 1 {
		return float64(quickSelect(nums, len(nums)/2))
	} else {
		return (float64(quickSelect(nums, len(nums)/2)) + float64(quickSelect(nums, len(nums)/2-1))) / 2
	}
}
