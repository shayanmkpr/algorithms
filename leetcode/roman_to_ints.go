package main

import "fmt"

func romanToInt(s string) int {
	// given a roman numeral give out the int
	fmt.Print(s)

	firstSingle := false
	var value int

	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := range s {
		char := string(s[len(s)-i-1])
		switch char {
		case "I", "X", "C", "M":
			if !firstSingle {
				value += romanMap[char]
			} else {
				value -= romanMap[char]
			}
		default:
			firstSingle = true
			value += romanMap[char]

		}
	}

	return value
}
