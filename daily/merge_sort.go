package main

/*
 the fact that you can append to fixed size array is interesting
*/

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left := mergeSort(arr[:len(arr)/2])
	right := mergeSort(arr[len(arr)/2:])
	return merge(left, right)
}

func merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	var ap, bp int
	for ap < len(a) && bp < len(b) {
		if a[ap] < b[bp] {
			result = append(result, a[ap])
			ap++
		} else {
			result = append(result, b[bp])
			bp++
		}
	}
	if ap <= len(a)-1 {
		result = append(result, a[ap:]...)
	} else {
		result = append(result, b[bp:]...)
	}
	return result
}
