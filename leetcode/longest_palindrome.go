package main

import "fmt"

func checkPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome0(s string) string {
	right := 1
	left := 0
	result := ""
	for right < len(s) && left < right {
		if checkPalindrome(s[left:right+1]) == true {
			fmt.Printf("yes")
			result = s[left : right+1]
			right += 2
			if left > 0 {
				left -= 1
			}
		} else if checkPalindrome(s[left:right]) == true {
			fmt.Printf("no")
			result = s[left:right]
			right += 1
			if left > 0 {
				left -= 1
			}
		} else {
			right += 1
			left += 1
		}
	}
	if len(s) == 1 {
		result = s
	}
	return result
}

func longestPalindrome(s string) string {
	l, r := 0, 1
	var result string
	for l < r && r < len(s) {
		palin := s[l:r]
		for i := 0; i <= (r-l)/2; i++ {
			if palin[i] == palin[len(palin)-i-1] {
				if i == (r-l)/2 {
					result = palin
					fmt.Println("0")
					r++
					if l > 0 {
						l--
					}
					break
				}
				continue
			} else if i != 0 && palin[i] == palin[len(palin)-i] {
				if i == (r-l)/2 {
					result = palin + string(s[r+1])
					fmt.Println("1")
					r += 2
					if l > 0 {
						l--
					}
					break
				}
				continue
			} else {
				l++
				r++
				break
			}
		}
	}
	return result
}
