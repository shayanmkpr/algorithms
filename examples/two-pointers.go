//go:build ignore

package examples

import "sort"

/*
================================================================
Algorithm: Two Pointers
================================================================

[EASY] LeetCode 125 - Valid Palindrome
Return true if the string is a palindrome considering only alphanumeric
characters and ignoring case.
Example: s = "A man, a plan, a canal: Panama"  ->  true
Hint: One pointer at each end. Skip non-alphanumerics. Compare lowercased
chars and move both inwards.

[EASY] LeetCode 26 - Remove Duplicates from Sorted Array
Remove duplicates in-place. Return the new length k; first k elements must
be unique in original order.
Example: nums = [1,1,2]  ->  k = 2, nums = [1,2,_]
Hint: Slow/fast pointers. `slow` is the write index for the next unique
value. Advance `fast` through the array; when nums[fast] != nums[slow-1],
write it to slow and increment slow.

[EASY] LeetCode 283 - Move Zeroes
Move all 0's to the end while maintaining the relative order of non-zero
elements. In-place.
Example: nums = [0,1,0,3,12]  ->  [1,3,12,0,0]
Hint: Slow pointer for write position. Walk fast through array; whenever
nums[fast] != 0, swap with nums[slow] and advance slow.

[MEDIUM] LeetCode 15 - 3Sum
Find all unique triplets in nums whose sum is 0.
Example: nums = [-1,0,1,2,-1,-4]  ->  [[-1,-1,2],[-1,0,1]]
Hint: Sort first. Fix index i, then use two pointers (l, r) on the rest.
Skip duplicates on i, l, and r to avoid repeated triplets.

[HARD] LeetCode 42 - Trapping Rain Water
Given an elevation map, compute how much water it can trap after rain.
Example: height = [0,1,0,2,1,0,1,3,2,1,2,1]  ->  6
Hint: Maintain leftMax and rightMax. Move the pointer on the smaller side
and add maxOnThatSide - height[pointer] to the answer.
*/

// ===================== Answers (Optimal Solutions) =====================

// 1) [EASY] LC 125: Valid Palindrome -- O(n) time, O(1) space
func isPalindrome(s string) bool {
	isAlnum := func(b byte) bool {
		return (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
	}
	toLower := func(b byte) byte {
		if b >= 'A' && b <= 'Z' {
			return b + 32
		}
		return b
	}
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isAlnum(s[l]) {
			l++
		}
		for l < r && !isAlnum(s[r]) {
			r--
		}
		if toLower(s[l]) != toLower(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}

// 2) [EASY] LC 26: Remove Duplicates from Sorted Array -- O(n)
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// 3) [EASY] LC 283: Move Zeroes -- O(n)
func moveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}

// 4) [MEDIUM] LC 15: 3Sum -- O(n^2) time, O(1) extra (excluding output)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			switch {
			case s == 0:
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			case s < 0:
				l++
			default:
				r--
			}
		}
	}
	return res
}

// 5) [HARD] LC 42: Trapping Rain Water -- O(n) time, O(1) space
func trap(height []int) int {
	l, r := 0, len(height)-1
	lMax, rMax, water := 0, 0, 0
	for l < r {
		if height[l] < height[r] {
			if height[l] >= lMax {
				lMax = height[l]
			} else {
				water += lMax - height[l]
			}
			l++
		} else {
			if height[r] >= rMax {
				rMax = height[r]
			} else {
				water += rMax - height[r]
			}
			r--
		}
	}
	return water
}
