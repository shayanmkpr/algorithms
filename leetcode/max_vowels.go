package main

func maxVowels(s string, k int) int {
	vowel := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	}
	var cnt, maxCnt int
	// make the window:
	for i := range k {
		if vowel[s[i]] {
			maxCnt++
		}
	}

	cnt = maxCnt
	for i := k; i < len(s); i++ {
		if vowel[s[i]] {
			cnt++
		}
		if vowel[s[i-k]] {
			cnt--
		}
		if cnt > maxCnt {
			maxCnt = cnt
		}
	}
	return maxCnt
}
