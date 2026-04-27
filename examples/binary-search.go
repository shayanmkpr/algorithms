//go:build ignore

package examples

/*
================================================================
Algorithm: Binary Search
================================================================

[EASY] LeetCode 704 - Binary Search
Given a sorted array `nums` and a target, return the index of target,
or -1 if not found. Must run in O(log n).
Example: nums = [-1,0,3,5,9,12], target = 9  ->  4
Hint: Keep `lo`, `hi`. Compute `mid = lo + (hi-lo)/2` to avoid overflow.

[EASY] LeetCode 35 - Search Insert Position
Return the index where target is, or where it should be inserted to keep
the array sorted.
Example: nums = [1,3,5,6], target = 2  ->  1
Hint: Search for the lower bound: when not found, `lo` ends up as the first
index with nums[i] >= target.

[EASY] LeetCode 278 - First Bad Version
Given an API isBadVersion(v) and n versions, find the first bad version
with minimum API calls.
Example: n = 5, first bad = 4  ->  4
Hint: Binary search the boundary. If isBadVersion(mid) -> hi = mid, else
lo = mid + 1. Stop when lo == hi.

[MEDIUM] LeetCode 33 - Search in Rotated Sorted Array
A sorted array was rotated at an unknown pivot. Find target in O(log n).
Example: nums = [4,5,6,7,0,1,2], target = 0  ->  4
Hint: At each step one half is always sorted. Decide which half holds the
target by comparing nums[lo], nums[mid], nums[hi].

[HARD] LeetCode 4 - Median of Two Sorted Arrays
Given two sorted arrays nums1 and nums2, return the median of the merged
sorted array in O(log(min(m, n))).
Example: nums1 = [1,3], nums2 = [2]  ->  2.0
Hint: Binary search the partition index of the smaller array so that
left side has (m+n+1)/2 elements and max(left) <= min(right).
*/

// ===================== Answers (Optimal Solutions) =====================

// 1) [EASY] LC 704: Binary Search  -- O(log n) time, O(1) space
func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return -1
}

// 2) [EASY] LC 35: Search Insert Position -- O(log n)
func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// 3) [EASY] LC 278: First Bad Version -- O(log n)
var isBadVersion = func(v int) bool { return false } // provided by judge

func firstBadVersion(n int) int {
	lo, hi := 1, n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if isBadVersion(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// 4) [MEDIUM] LC 33: Search in Rotated Sorted Array -- O(log n)
func searchRotated(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		if nums[lo] <= nums[mid] { // left half sorted
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else { // right half sorted
			if nums[mid] < target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

// 5) [HARD] LC 4: Median of Two Sorted Arrays -- O(log(min(m, n)))
func findMedianSortedArrays(nums1, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	lo, hi, half := 0, m, (m+n+1)/2
	const inf = int(1<<31 - 1)
	for lo <= hi {
		i := (lo + hi) / 2
		j := half - i
		l1, r1 := -inf, inf
		l2, r2 := -inf, inf
		if i > 0 {
			l1 = nums1[i-1]
		}
		if i < m {
			r1 = nums1[i]
		}
		if j > 0 {
			l2 = nums2[j-1]
		}
		if j < n {
			r2 = nums2[j]
		}
		if l1 <= r2 && l2 <= r1 {
			if (m+n)%2 == 1 {
				return float64(max(l1, l2))
			}
			return float64(max(l1, l2)+min(r1, r2)) / 2.0
		} else if l1 > r2 {
			hi = i - 1
		} else {
			lo = i + 1
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
