//go:build ignore

package examples

import "math/rand"

/*
================================================================
Algorithm: Quick Sort (and Quickselect)
================================================================

[EASY] LeetCode 912 - Sort an Array
Sort the array in O(n log n) on average. Implement quick sort.
Example: nums = [5,2,3,1]  ->  [1,2,3,5]
Hint: Pick a random pivot to avoid worst-case O(n^2). Partition into
< pivot and > pivot, then recurse.

[EASY] LeetCode 75 - Sort Colors (Dutch National Flag)
Given an array with values 0, 1, 2, sort them in-place in one pass.
Example: nums = [2,0,2,1,1,0]  ->  [0,0,1,1,2,2]
Hint: Three-way partitioning. Keep `lo` (next 0 slot), `hi` (next 2 slot),
`i` (current). On 0 swap to lo and advance both; on 2 swap to hi and shrink;
on 1 just advance i.

[EASY] LeetCode 905 - Sort Array By Parity
Move all even numbers before all odd numbers (any internal order).
Example: nums = [3,1,2,4]  ->  [2,4,3,1]  (one valid answer)
Hint: Two-pointer partition (the heart of quick sort). i goes from left
swapping evens, j from right; swap when nums[i] is odd and nums[j] is
even, otherwise advance.

[MEDIUM] LeetCode 215 - Kth Largest Element in an Array
Return the kth largest element. Average O(n) with quickselect.
Example: nums = [3,2,1,5,6,4], k = 2  ->  5
Hint: Convert to "find element with index n-k after partition". Recurse
into only one side after each partition.

[HARD] LeetCode 324 - Wiggle Sort II
Reorder nums so nums[0] < nums[1] > nums[2] < nums[3] ... in O(n) average
time, ideally O(1) extra space.
Example: nums = [1,5,1,1,6,4]  ->  [1,6,1,5,1,4] (one valid answer)
Hint: Quickselect to find the median. Then place smaller half on even
indexes and larger half on odd indexes (use virtual indexing trick).
*/

// ===================== Answers (Optimal Solutions) =====================

// 1) [EASY] LC 912: Sort an Array -- average O(n log n)
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func partition(a []int, lo, hi int) int {
	r := lo + rand.Intn(hi-lo+1)
	a[r], a[hi] = a[hi], a[r]
	pivot := a[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}

// 2) [EASY] LC 75: Sort Colors -- O(n) one-pass, O(1) space
func sortColors(nums []int) {
	lo, hi, i := 0, len(nums)-1, 0
	for i <= hi {
		switch nums[i] {
		case 0:
			nums[i], nums[lo] = nums[lo], nums[i]
			lo++
			i++
		case 2:
			nums[i], nums[hi] = nums[hi], nums[i]
			hi--
		default:
			i++
		}
	}
}

// 3) [EASY] LC 905: Sort Array By Parity -- O(n) one-pass, O(1) space
func sortArrayByParity(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]%2 > nums[j]%2 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		if nums[i]%2 == 0 {
			i++
		}
		if nums[j]%2 == 1 {
			j--
		}
	}
	return nums
}

// 4) [MEDIUM] LC 215: Kth Largest Element -- average O(n) via Quickselect
func findKthLargest(nums []int, k int) int {
	target := len(nums) - k
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		p := partition(nums, lo, hi)
		switch {
		case p == target:
			return nums[p]
		case p < target:
			lo = p + 1
		default:
			hi = p - 1
		}
	}
	return -1
}

// 5) [HARD] LC 324: Wiggle Sort II -- O(n) avg via quickselect + virtual index
func wiggleSort(nums []int) {
	n := len(nums)
	mid := quickSelect(append([]int(nil), nums...), 0, n-1, n/2)

	idx := func(i int) int { return (1 + 2*i) % (n | 1) }

	i, j, k := 0, n-1, 0
	for k <= j {
		if nums[idx(k)] > mid {
			nums[idx(i)], nums[idx(k)] = nums[idx(k)], nums[idx(i)]
			i++
			k++
		} else if nums[idx(k)] < mid {
			nums[idx(k)], nums[idx(j)] = nums[idx(j)], nums[idx(k)]
			j--
		} else {
			k++
		}
	}
}

func quickSelect(a []int, lo, hi, k int) int {
	for lo <= hi {
		p := partition(a, lo, hi)
		switch {
		case p == k:
			return a[p]
		case p < k:
			lo = p + 1
		default:
			hi = p - 1
		}
	}
	return a[lo]
}
