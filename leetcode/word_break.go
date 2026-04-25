package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	fmt.Println(s)
	fmt.Println(wordDict)
	// check for the reference in the s string.
	var rec func(s string, wordDict []string) bool
	rec = func(s string, wordDict []string) bool {
		i := 1
		for i <= len(s) {
			for _, word := range wordDict {
				if s[:i] == word {
					if len(s) == 0 { // we won
						fmt.Println("we won")
						return true
					} else {
						fmt.Println("yes")
						fmt.Println(s[:i])
						return rec(s[i:], wordDict)
					}
				} else {
					fmt.Println("###################")
					fmt.Println(s[:i])
				}
			}
			i++
		}
		return false
	}
	return rec(s, wordDict)
}
