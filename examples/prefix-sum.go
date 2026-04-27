//go:build ignore

package examples

import "sort"

/*
================================================================
Algorithm: Prefix Sum
================================================================

[EASY] LeetCode 303 - Range Sum Query - Immutable
Implement NumArray supporting sumRange(l, r) returning sum of nums[l..r].
Example: nums=[-2,0,3,-5,2,-1]; sumRange(0,2)->1; sumRange(2,5)->-1
Hint: Precompute prefix[i] = nums[0]+...+nums[i-1] in O(n). Each query is
prefix[r+1] - prefix[l] in O(1).

[EASY] LeetCode 1480 - Running Sum of 1d Array
Return result[i] = nums[0] + nums[1] + ... + nums[i].
Example: nums = [1,2,3,4]  ->  [1,3,6,10]
Hint: This is the prefix sum itself. Iterate once, maintain a running sum,
write into result[i] (or in-place into nums).

[EASY] LeetCode 724 - Find Pivot Index
Return the leftmost index i where sum(nums[:i]) == sum(nums[i+1:]); else -1.
Example: nums = [1,7,3,6,5,6]  ->  3
Hint: Compute total sum. Walk left-to-right tracking left sum; right sum =
total - left - nums[i]. Return when they're equal.

[MEDIUM] LeetCode 560 - Subarray Sum Equals K
Return number of contiguous subarrays whose sum equals k.
Example: nums = [1,1,1], k = 2  ->  2
Hint: Track running prefix `sum`. For each index, count how many earlier
prefix values equal sum-k using a hash map (init {0:1}).

[HARD] LeetCode 363 - Max Sum of Rectangle No Larger Than K
Given an m x n matrix, find the max sum of a rectangle whose sum <= k.
Example: matrix=[[1,0,1],[0,-2,3]], k=2  ->  2
Hint: Fix two columns (or rows). Build row-wise sums between them, then
for each prefix sum search the smallest prefix > prefix-k using a sorted
structure (Fenwick / TreeSet / sorted slice + binary search).
*/

// ===================== Answers (Optimal Solutions) =====================

// 1) [EASY] LC 303: Range Sum Query - Immutable -- O(n) build, O(1) query
type NumArray struct {
	prefix []int
}

func ConstructorNumArray(nums []int) NumArray {
	p := make([]int, len(nums)+1)
	for i, v := range nums {
		p[i+1] = p[i] + v
	}
	return NumArray{prefix: p}
}

func (na *NumArray) SumRange(left, right int) int {
	return na.prefix[right+1] - na.prefix[left]
}

// 2) [EASY] LC 1480: Running Sum of 1d Array -- O(n)
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

// 3) [EASY] LC 724: Find Pivot Index -- O(n)
func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	left := 0
	for i, v := range nums {
		if left == total-left-v {
			return i
		}
		left += v
	}
	return -1
}

// 4) [MEDIUM] LC 560: Subarray Sum Equals K -- O(n) time, O(n) space
func subarraySum(nums []int, k int) int {
	count, sum := 0, 0
	seen := map[int]int{0: 1}
	for _, v := range nums {
		sum += v
		count += seen[sum-k]
		seen[sum]++
	}
	return count
}

// 5) [HARD] LC 363: Max Sum of Rectangle No Larger Than K -- O(min(m,n)^2 * max(m,n) * log(max(m,n)))
func maxSumSubmatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	const negInf = -1 << 31
	best := negInf

	rows, cols := m, n
	transpose := false
	if cols < rows { // iterate over the smaller dimension as "left/right"
		rows, cols = cols, rows
		transpose = true
	}

	for left := 0; left < rows; left++ {
		sums := make([]int, cols)
		for right := left; right < rows; right++ {
			for j := 0; j < cols; j++ {
				if transpose {
					sums[j] += matrix[j][right]
				} else {
					sums[j] += matrix[right][j]
				}
			}
			// find max subarray sum <= k in `sums`
			prefixes := []int{0}
			cur := 0
			for _, v := range sums {
				cur += v
				// smallest prefix >= cur-k
				idx := sort.SearchInts(prefixes, cur-k)
				if idx < len(prefixes) {
					if cur-prefixes[idx] > best {
						best = cur - prefixes[idx]
					}
				}
				ins := sort.SearchInts(prefixes, cur)
				prefixes = append(prefixes, 0)
				copy(prefixes[ins+1:], prefixes[ins:])
				prefixes[ins] = cur
			}
		}
	}
	return best
}
