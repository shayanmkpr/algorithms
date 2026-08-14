package main

// func myQuickSort(arr []int) {
// 	if len(arr) <= 1 {
// 		return
// 	}
//
// 	pivot := (len(arr) - 1) / 2
// 	var left, right, equal []int
// 	for _, v := range arr {
// 		if v < arr[pivot] {
// 			left = append(left, v)
// 		}
// 		if v > arr[pivot] {
// 			right = append(right, v)
// 		}
// 		if v == arr[pivot] {
// 			equal = append(equal, v)
// 		}
// 	}
// 	myQuickSort(left)
// 	myQuickSort(right)
// 	copy(arr, append(append(left, equal...), right...))
// }

type stuff struct {
	someString string
}

type item struct {
	k     int
	value stuff
}

func myQuickSort(arr []item) { // inplace sort
	if len(arr) == 0 {
		return
	}
	pivot := len(arr) / 2 // incase half sorted arr, tc = o(n^2)
	var left, eq, right []item
	for i := 0; i < len(arr); i++ {
		if arr[i].k < arr[pivot].k {
			left = append(left, arr[i])
		}
		if arr[i].k == arr[pivot].k {
			eq = append(eq, arr[i])
		}
		if arr[i].k > arr[pivot].k {
			right = append(right, arr[i])
		}
	}
	myQuickSort(left)
	myQuickSort(right)
	copy(arr, append(append(left, eq...), right...))
}
