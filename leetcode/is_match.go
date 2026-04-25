package main

import "fmt"

func isMatch(s string, p string) bool {
	// given a string and a pattern check if the string matches the pattern.

	type starStruct struct {
		Pointer int
		Flag    bool
	}

	if len(p) == 0 {
		return len(s) == 0
	}

	starPointer := starStruct{
		Pointer: 0,
		Flag:    false,
	}

	// avoid panic when checking the last char
	if len(p) > 0 && p[len(p)-1] != '*' {
		if len(s) < len(p) {
			fmt.Println("string is shorter than the pattern")
			return false
		}
	} else if len(p) > 1 {
		if len(s) < len(p)-1 {
			fmt.Println("string is shorter than the pattern")
			return false
		}
	}

	for ip := 0; ip < len(p) && ip < len(s); ip++ {
		charp := p[ip]

		if charp == s[ip] {
			continue
		} else if charp == '.' {
			continue
		} else if charp == '*' {
			// everything from now on should be the same as the char before *
			if !starPointer.Flag {
				starPointer.Pointer = ip
				starPointer.Flag = true
			}
			continue
		} else if starPointer.Flag {
			if s[ip] == p[starPointer.Pointer-1] {
				continue
			}
		} else {
			return false
		}
	}
	// after your loop
	if len(s) != len(p) {
		return false
	}
	return true
}
