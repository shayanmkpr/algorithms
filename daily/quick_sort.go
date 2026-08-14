package main

func quickSort(arr []int) []int { // in place quick sort
	// you find a pivot, you divide the array into 3 parts, less, equal, and more.
	// then you do the same on each part
	if len(arr) < 1 {
		return []int{}
	}
	pivot := (len(arr) - 1) / 2
	var less, equal, more []int
	for _, v := range arr {
		if v < arr[pivot] {
			less = append(less, v)
		}
		if v > arr[pivot] {
			more = append(more, v)
		}
		if v == arr[pivot] {
			equal = append(equal, v)
		}
	}
	less = quickSort(less)
	more = quickSort(more)
	// copy(arr, append(less, append(equal, more...)...))
	// return arr
	return append(less, append(equal, more...)...)
}
