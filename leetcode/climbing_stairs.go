package main

func climbStairs(n int) int {
	mem := map[int]int{1: 1, 0: 1}
	var rec func(r int) int
	rec = func(r int) int {
		if _, ok := mem[r]; ok {
			return mem[r]
		}
		mem[r] = rec(r-2) + rec(r-1)
		return mem[r]
	}
	return rec(n)
}
