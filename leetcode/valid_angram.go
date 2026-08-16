package main

func isAnagram(s string, t string) bool {
	seen := map[rune]int{}
	for _, ch := range s {
		seen[ch]++
	}

	for _, ch := range t {
		if _, ok := seen[ch]; ok {
			seen[ch]--
		} else {
			return false
		}
	}

	for _, v := range seen {
		if v != 0 {
			return false
		}
	}
	return true
}
