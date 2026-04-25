package main

func climbStairs(n int) int {
	memo := make(map[int]int)

	var rec func(int) int

	rec = func(n int) int {
		if n == 0 || n == 1 {
			return 1
		}
		// check if you have it
		if _, ok := memo[n]; ok {
			return memo[n]
		} else {
			memo[n] = rec(n-2) + rec(n-1)
			return memo[n]
		}
	}

	return rec(n)
}
