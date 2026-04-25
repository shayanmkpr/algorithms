package main

import "fmt"

// QuickSort - Average O(n log n), Worst O(n²)
// Divide and conquer: pick pivot, partition around it, recurse on both sides
func quickSort(arr []int) []int {
	// Base case: arrays with 0 or 1 element are already sorted
	if len(arr) < 2 {
		return arr
	}

	// Choose first element as pivot (could use random for better performance)
	pivot := arr[0]
	var less, greater []int

	// Partition: split elements into less than and greater than pivot
	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	// Recursively sort both partitions and combine with pivot
	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)
	return result
}

// MergeSort - Consistent O(n log n), O(n) space
// Divide array in half, sort each half, then merge sorted halves
func mergeSort(arr []int) []int {
	// Base case: single elements are already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Divide: split array in middle
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])  // Sort left half
	right := mergeSort(arr[mid:]) // Sort right half

	// Conquer: merge the two sorted halves
	var result []int
	i, j := 0, 0

	// Compare elements from both halves and pick smaller one
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Add remaining elements (one array will be exhausted first)
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// HeapSort - Consistent O(n log n), O(1) space (in-place)
// Build max heap, then repeatedly extract maximum and rebuild heap
func heapSort(arr []int) {
	n := len(arr)

	// Build max heap: start from last non-leaf node and heapify down
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		// Move current root (maximum) to end of array
		arr[0], arr[i] = arr[i], arr[0]

		// Reduce heap size and heapify root again
		heapify(arr, i, 0)
	}
}

// heapify maintains max heap property: parent >= children
func heapify(arr []int, heapSize, rootIndex int) {
	largest := rootIndex
	left := 2*rootIndex + 1  // Left child index
	right := 2*rootIndex + 2 // Right child index

	// Find largest among root and its children
	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}
	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}

	// If largest is not root, swap and continue heapifying
	if largest != rootIndex {
		arr[rootIndex], arr[largest] = arr[largest], arr[rootIndex]
		heapify(arr, heapSize, largest) // Recursively heapify affected subtree
	}
}

// IntroSort - Hybrid O(n log n) guaranteed
// Start with quicksort, switch to heapsort if recursion too deep, use insertion for small arrays
func introSort(arr []int) {
	// Calculate maximum recursion depth (2 * log2(n))
	maxDepth := 0
	for n := len(arr); n > 1; n /= 2 {
		maxDepth++
	}
	maxDepth *= 2

	introSortHelper(arr, 0, len(arr)-1, maxDepth)
}

func introSortHelper(arr []int, low, high, maxDepth int) {
	size := high - low + 1

	// Use insertion sort for small arrays (more efficient for small n)
	if size <= 16 {
		insertionSort(arr, low, high)
		return
	}

	// Use heap sort if recursion depth limit reached (avoid O(n²) worst case)
	if maxDepth == 0 {
		heapSortRange(arr, low, high)
		return
	}

	// Use quicksort for medium arrays
	pivotIndex := partition(arr, low, high)
	introSortHelper(arr, low, pivotIndex-1, maxDepth-1)
	introSortHelper(arr, pivotIndex+1, high, maxDepth-1)
}

// Insertion sort for small arrays in IntroSort
func insertionSort(arr []int, low, high int) {
	for i := low + 1; i <= high; i++ {
		key := arr[i]
		j := i - 1
		// Shift elements greater than key to the right
		for j >= low && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Heap sort for a range (used in IntroSort)
func heapSortRange(arr []int, low, high int) {
	// Extract the range, sort it, then put it back
	temp := make([]int, high-low+1)
	copy(temp, arr[low:high+1])
	heapSort(temp)
	copy(arr[low:high+1], temp)
}

// Partition function for IntroSort (same as quicksort partition)
func partition(arr []int, low, high int) int {
	pivot := arr[high] // Choose last element as pivot
	i := low - 1       // Index of smaller element

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // Place pivot in correct position
	return i + 1
}

// TimSort - Adaptive O(n log n), excellent for partially sorted data
// Find natural runs, extend short runs with insertion sort, then merge runs
func timSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	minRun := getMinRunLength(len(arr)) // Optimal run size (between 32-64)
	runs := [][]int{}                   // Store all runs

	i := 0
	for i < len(arr) {
		// Find existing run (ascending or descending)
		runStart := i
		if i == len(arr)-1 {
			runs = append(runs, []int{arr[i]})
			break
		}

		// Check if run is ascending or descending
		if arr[i] <= arr[i+1] {
			// Ascending run
			for i < len(arr)-1 && arr[i] <= arr[i+1] {
				i++
			}
		} else {
			// Descending run - reverse it to make ascending
			for i < len(arr)-1 && arr[i] > arr[i+1] {
				i++
			}
			// Reverse the descending run
			for left, right := runStart, i; left < right; left, right = left+1, right-1 {
				arr[left], arr[right] = arr[right], arr[left]
			}
		}
		i++

		// Extract the run
		run := make([]int, i-runStart)
		copy(run, arr[runStart:i])

		// Extend short runs with insertion sort to reach minRun length
		if len(run) < minRun && i < len(arr) {
			// Add more elements to reach minRun size
			need := minRun - len(run)
			if i+need > len(arr) {
				need = len(arr) - i
			}
			for j := 0; j < need; j++ {
				run = append(run, arr[i+j])
			}
			i += need

			// Sort the extended run with insertion sort
			for j := len(run) - need; j < len(run); j++ {
				key := run[j]
				k := j - 1
				for k >= 0 && run[k] > key {
					run[k+1] = run[k]
					k--
				}
				run[k+1] = key
			}
		}

		runs = append(runs, run)
	}

	// Merge all runs together
	for len(runs) > 1 {
		newRuns := [][]int{}
		for i := 0; i < len(runs); i += 2 {
			if i+1 < len(runs) {
				// Merge two runs
				merged := mergeRuns(runs[i], runs[i+1])
				newRuns = append(newRuns, merged)
			} else {
				// Odd run out, just add it
				newRuns = append(newRuns, runs[i])
			}
		}
		runs = newRuns
	}

	return runs[0]
}

// Calculate minimum run length for TimSort (between 32-64)
func getMinRunLength(n int) int {
	r := 0
	for n >= 32 {
		r |= n & 1 // Keep track if we had odd numbers
		n >>= 1    // Divide by 2
	}
	return n + r // Result will be between 32-64
}

// Merge two sorted runs for TimSort
func mergeRuns(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Merge in sorted order
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Add remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// Test and demonstration
func notMain() {
	// Test array
	original := []int{64, 34, 25, 12, 22, 11, 90, 5, 77, 30}

	fmt.Println("Original array:", original)
	fmt.Println()

	// Test QuickSort
	arr1 := make([]int, len(original))
	copy(arr1, original)
	sorted1 := quickSort(arr1)
	fmt.Println("QuickSort:     ", sorted1)

	// Test MergeSort
	arr2 := make([]int, len(original))
	copy(arr2, original)
	sorted2 := mergeSort(arr2)
	fmt.Println("MergeSort:     ", sorted2)

	// Test HeapSort (modifies original array)
	arr3 := make([]int, len(original))
	copy(arr3, original)
	heapSort(arr3)
	fmt.Println("HeapSort:      ", arr3)

	// Test IntroSort (modifies original array)
	arr4 := make([]int, len(original))
	copy(arr4, original)
	introSort(arr4)
	fmt.Println("IntroSort:     ", arr4)

	// Test TimSort
	arr5 := make([]int, len(original))
	copy(arr5, original)
	sorted5 := timSort(arr5)
	fmt.Println("TimSort:       ", sorted5)
}
