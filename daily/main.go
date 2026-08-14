package main

import "fmt"

type loc struct {
	x int
	y int
}

type grid map[loc][]loc

type value = int

func main() {
	someArr := []int{1, 5, 3, 2, 3, 4, 5, 6, 6, 3, 4, 3}
	quickSort(someArr)
	fmt.Println(someArr)
	fmt.Println(quickSort(someArr))
	fmt.Println(countNumbersWithUniqueDigits(3))
	// fmt.Println(mergeSort(someArr))
	// fmt.Println(pascalSecond(3))
	// fmt.Println(dq("(()((())))"))
}

func runBfs() {
	myGrid := grid{
		loc{x: 0, y: 0}: []loc{{x: 0, y: 1}, {x: 1, y: 0}},
		loc{x: 0, y: 1}: []loc{{x: 0, y: 0}, {x: 0, y: 2}, {x: 1, y: 1}},
		loc{x: 0, y: 2}: []loc{{x: 0, y: 1}, {x: 1, y: 2}},
		loc{x: 1, y: 0}: []loc{{x: 0, y: 0}, {x: 1, y: 1}, {x: 2, y: 0}},
		loc{x: 1, y: 1}: []loc{{x: 0, y: 1}, {x: 1, y: 0}, {x: 1, y: 2}, {x: 2, y: 1}},
		loc{x: 1, y: 2}: []loc{{x: 0, y: 2}, {x: 1, y: 1}, {x: 2, y: 2}},
		loc{x: 2, y: 0}: []loc{{x: 1, y: 0}, {x: 2, y: 1}},
		loc{x: 2, y: 1}: []loc{{x: 2, y: 0}, {x: 1, y: 1}, {x: 2, y: 2}},
		loc{x: 2, y: 2}: []loc{{x: 1, y: 2}, {x: 2, y: 1}},
	}
	_ = myGrid
}
