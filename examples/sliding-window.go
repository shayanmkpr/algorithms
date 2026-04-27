//go:build ignore

package examples

/*
================================================================
Algorithm: Sliding Window
================================================================

[EASY] LeetCode 643 - Maximum Average Subarray I
Find the contiguous subarray of length k that has the maximum average.
Example: nums = [1,12,-5,-6,50,3], k = 4  ->  12.75
Hint: Compute sum of first k elements, then slide: add new, drop old.
Track max sum.

[EASY] LeetCode 219 - Contains Duplicate II
Return true if there exist i != j with nums[i] == nums[j] and |i-j| <= k.
Example: nums = [1,2,3,1], k = 3  ->  true
Hint: Maintain a set/window of the last k elements. If nums[i] is already
in it, return true; else add and evict nums[i-k] when i >= k.

[EASY] LeetCode 1876 - Substrings of Size Three with Distinct Characters
Count substrings of length 3 with all distinct characters.
Example: s = "xyzzaz"  ->  1 ("xyz")
Hint: Fixed-size window of length 3. As you slide, check if all 3 chars
are distinct; increment counter on success.

[MEDIUM] LeetCode 3 - Longest Substring Without Repeating Characters
Given a string s, find the length of the longest substring without
repeating characters.
Example: s = "abcabcbb"  ->  3 ("abc")
Hint: Expand right, store last index of each char in a map. If char repeats
inside window, move left to lastIndex+1. Update max length each step.

[HARD] LeetCode 76 - Minimum Window Substring
Given strings s and t, return the minimum window in s that contains all
characters of t (with multiplicity). If none, return "".
Example: s = "ADOBECODEBANC", t = "ABC"  ->  "BANC"
Hint: Use a freq map of t and a `have/need` counter. Expand right; when
have == need, shrink left while window is still valid and update answer.
*/

// ===================== Answers (Optimal Solutions) =====================

// 1) [EASY] LC 643: Maximum Average Subarray I -- O(n)
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	best := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > best {
			best = sum
		}
	}
	return float64(best) / float64(k)
}

// 2) [EASY] LC 219: Contains Duplicate II -- O(n) time, O(k) space
func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]struct{}, k)
	for i, v := range nums {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = struct{}{}
		if i >= k {
			delete(seen, nums[i-k])
		}
	}
	return false
}

// 3) [EASY] LC 1876: Substrings of Size Three with Distinct Characters -- O(n)
func countGoodSubstrings(s string) int {
	if len(s) < 3 {
		return 0
	}
	count := 0
	for i := 0; i+3 <= len(s); i++ {
		if s[i] != s[i+1] && s[i] != s[i+2] && s[i+1] != s[i+2] {
			count++
		}
	}
	return count
}

// 4) [MEDIUM] LC 3: Longest Substring Without Repeating Characters -- O(n)
func lengthOfLongestSubstring(s string) int {
	last := make(map[byte]int)
	left, best := 0, 0
	for right := 0; right < len(s); right++ {
		if idx, ok := last[s[right]]; ok && idx >= left {
			left = idx + 1
		}
		last[s[right]] = right
		if right-left+1 > best {
			best = right - left + 1
		}
	}
	return best
}

// 5) [HARD] LC 76: Minimum Window Substring -- O(|s| + |t|)
func minWindow(s, t string) string {
	if len(t) > len(s) {
		return ""
	}
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	required := len(need)

	have := 0
	window := make(map[byte]int)
	bestL, bestLen := 0, -1
	left := 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++
		if window[c] == need[c] {
			have++
		}
		for have == required {
			if bestLen == -1 || right-left+1 < bestLen {
				bestL, bestLen = left, right-left+1
			}
			lc := s[left]
			window[lc]--
			if window[lc] < need[lc] {
				have--
			}
			left++
		}
	}
	if bestLen == -1 {
		return ""
	}
	return s[bestL : bestL+bestLen]
}
