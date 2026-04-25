package main

import "fmt"

func maxArea(height []int) int {
	var p1 int
	var p2 int
	var area int
	var temp int
	p1, p2 = 0, len(height)-1
	area, temp = 0, 0
	fmt.Println(height)

	for range len(height) {
		if height[p1] < height[p2] {
			temp = height[p1] * (p2 - p1)
			for i := 1; p1+i < len(height)-1 && p1+i < p2; i++ {
				if height[p1+i] > height[p1] {
					p1 += i
					break
				}
			}
		} else {
			temp = height[p2] * (p2 - p1)
			for i := 1; p2-i > 0 && p2-i > p1; i++ {
				if height[p2-i] > height[p2] {
					p2 -= i
					break
				}
			}
		}
		if temp >= area {
			area = temp
		}
	}
	return area
}
