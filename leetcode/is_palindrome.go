package main

import "fmt"

func reverse(x int) int {
	var result int = 0
	var negative bool

	if x < 0 {
		negative = true
		x = -1 * x
	} else {
		negative = false
	}

	var div int = x

	for div >= 1 {
		rem := div % 10
		div = div / 10
		result = result*10 + rem
	}

	if result < -2147483648 || result > 2147483647 {
		return 0
	} else if !negative {
		return result
	} else if negative {
		return -1 * result
	} else {
		return 0
	}
}

func isPalindrome(x int) bool {
	// check if the x is negative:
	if x < 0 {
		fmt.Println("negative input \n")
		return false
	}

	var div int = x
	var rem int = 0
	var reverse int = 0

	for div != 0 {
		rem = div % 10
		div = div / 10
		reverse = reverse*10 + rem
	}

	if reverse == x {
		return true
	}

	return false
}
