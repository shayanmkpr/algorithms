package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	myMap := make(map[int]bool)
	var length int = 0
	lengthMax := length
	var curr int

	for _, v := range nums {
		myMap[v] = true
	}

	for value := range myMap {
		if !myMap[value-1] {
			length = 1
			curr = value
		}

		for myMap[curr+1] {
			fmt.Println(curr + 1)
			curr++
			length++
		}

		if length > lengthMax {
			lengthMax = length
		}
	}

	return lengthMax
}
