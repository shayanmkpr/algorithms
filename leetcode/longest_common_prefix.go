package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	prf := strs[0]
	for _, word := range strs {
		for !strings.HasPrefix(word, prf) {
			prf = prf[:len(prf)-1]
		}
	}
	return prf
}
