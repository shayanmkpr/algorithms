package main

func lengthOfLongestSubstring(s string) int {
	max := 0
	left := 0
	seen := make(map[byte]int)
	for right := 0; right < len(s); right++ {
		if i, ok := seen[s[right]]; ok && i >= left {
			left = i + 1
		}
		if right-left+1 > max {
			max = right - left + 1
		}
		seen[s[right]] = right
	}
	return max
}
