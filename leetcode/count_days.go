package main

import (
	"fmt"
	"sort"
)

func countDays(days int, meetings [][]int) int {
	fmt.Println("input", meetings, days)
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	var pe, off int
	for _, m := range meetings {
		s, e := m[0], m[1]
		if s > pe {
			off += s - pe - 1
		}
		pe = e
	}
	return off
}
