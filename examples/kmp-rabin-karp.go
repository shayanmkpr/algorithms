//go:build ignore

package examples

/*
================================================================
Algorithm: KMP / Rabin-Karp (String Matching)
================================================================

[EASY] LeetCode 28 - Find the Index of the First Occurrence in a String
Return the index of the first occurrence of needle in haystack, or -1.
Example: haystack = "sadbutsad", needle = "sad"  ->  0
Hint: Build the KMP prefix (longest-proper-prefix/suffix) table for
needle, then scan haystack; on mismatch jump j = lps[j-1].

[EASY] LeetCode 459 - Repeated Substring Pattern
Return true if s can be constructed by repeating a substring multiple
times.
Example: s = "abab"  ->  true ("ab" repeated twice)
Hint: Build the KMP lps array of s. If lps[n-1] > 0 and n % (n -
lps[n-1]) == 0, the smallest repeating unit has length n - lps[n-1].

[EASY] LeetCode 796 - Rotate String
Return true if s can become goal after some number of shifts.
Example: s = "abcde", goal = "cdeab"  ->  true
Hint: If lengths differ, false. Else check whether goal is a substring
of s+s using KMP substring search.

[MEDIUM] LeetCode 1392 - Longest Happy Prefix
Return the longest proper prefix of s which is also a suffix.
Example: s = "level"  ->  "l"   ("l" is both a prefix and a suffix)
Hint: Compute the KMP lps array; the answer is s[:lps[n-1]].

[MEDIUM] LeetCode 686 - Repeated String Match
Return the minimum number of times string a must be repeated so that b
becomes a substring of it. If impossible, return -1.
Example: a = "abcd", b = "cdabcdab"  ->  3
Hint: Repeat a at least ceil(len(b)/len(a)) times; also try one more
copy. Use KMP (strStr) to check whether b is contained. If the index
lands beyond the repeated region, return -1.

[HARD] LeetCode 1044 - Longest Duplicate Substring
Return any longest substring that occurs at least twice in s.
Example: s = "banana"  ->  "ana"
Hint: Binary search the length L; for each L use Rabin-Karp rolling
hash with a large modulus (and a second hash to avoid collisions) to
find a duplicate; if found, try a longer L.
*/

// ===================== Answers (Optimal Solutions) =====================

// 1) [EASY] LC 28: strStr (KMP) -- O(n + m)
func strStr(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	lps := buildLPS(needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = lps[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}

// buildLPS returns the longest-proper-prefix-suffix table for KMP.
func buildLPS(s string) []int {
	lps := make([]int, len(s))
	length := 0
	for i := 1; i < len(s); i++ {
		for length > 0 && s[i] != s[length] {
			length = lps[length-1]
		}
		if s[i] == s[length] {
			length++
		}
		lps[i] = length
	}
	return lps
}

// 2) [EASY] LC 459: Repeated Substring Pattern -- O(n)
func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	lps := buildLPS(s)
	p := lps[n-1]
	return p > 0 && n%(n-p) == 0
}

// 3) [EASY] LC 796: Rotate String -- O(n) via KMP
func rotateString(s, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	return strStr(s+s, goal) != -1
}

// 4) [MEDIUM] LC 1392: Longest Happy Prefix -- O(n)
func longestPrefix(s string) string {
	lps := buildLPS(s)
	return s[:lps[len(s)-1]]
}

// 5) [MEDIUM] LC 686: Repeated String Match -- O((n + m) * n/m)
func repeatedStringMatch(a, b string) int {
	// Repeat a just enough to cover b (and one extra for a partial wrap).
	repeats := (len(b) + len(a) - 1) / len(a)
	sb := a
	for i := 1; i < repeats; i++ {
		sb += a
	}
	if idx := strStr(sb, b); idx != -1 {
		return repeats
	}
	sb += a
	if idx := strStr(sb, b); idx != -1 && idx < repeats*len(a) {
		return repeats + 1
	}
	return -1
}

// 6) [HARD] LC 1044: Longest Duplicate Substring -- O(n log n) avg
func longestDupSubstring(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	const base = 26
	const mod1 = int(1e9 + 7)
	const mod2 = int(1e9 + 9)

	pow1 := make([]int, n+1)
	pow2 := make([]int, n+1)
	pow1[0], pow2[0] = 1, 1
	for i := 1; i <= n; i++ {
		pow1[i] = (pow1[i-1] * base) % mod1
		pow2[i] = (pow2[i-1] * base) % mod2
	}

	// Returns the start index of a duplicate substring of length L, or -1.
	findDup := func(L int) int {
		h1, h2 := 0, 0
		for i := 0; i < L; i++ {
			v := int(s[i] - 'a')
			h1 = (h1*base + v) % mod1
			h2 = (h2*base + v) % mod2
		}
		seen := make(map[[2]int]int)
		seen[[2]int{h1, h2}] = 0
		for i := L; i < n; i++ {
			v := int(s[i] - 'a')
			out := int(s[i-L] - 'a')
			h1 = (h1*base + v - out*pow1[L]%mod1 + mod1*pow1[L]) % mod1
			h2 = (h2*base + v - out*pow2[L]%mod2 + mod2*pow2[L]) % mod2
			key := [2]int{h1, h2}
			if idx, ok := seen[key]; ok {
				if s[idx:idx+L] == s[i-L+1:i+1] {
					return idx
				}
			} else {
				seen[key] = i - L + 1
			}
		}
		return -1
	}

	lo, hi := 1, n-1
	start, length := 0, 0
	for lo <= hi {
		mid := (lo + hi) / 2
		idx := findDup(mid)
		if idx != -1 {
			start, length = idx, mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return s[start : start+length]
}
