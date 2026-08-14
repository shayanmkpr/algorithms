package main

import "math"

func dq(s string) float64 {
	if s == "" {
		return 0.0
	}
	var cnt int
	var node float64
	var temp string
	for _, ch := range s {
		if ch == '(' {
			cnt++
		}
		if ch == ')' {
			cnt--
		}
		temp += string(ch)
		if cnt == 0 {
			node += math.Pow(2, dq(temp[1:len(temp)-1]))
		}
	}
	return node
}

func wiggleSort(nums []int) {
}
