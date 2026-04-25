package main

func numTrees(n int) int {
	mem := make(map[int]int) // parrent --> number of ways
	mem[1] = 1
	mem[0] = 1

	var rec func(n int) int
	rec = func(n int) int {
		if value, ok := mem[n]; ok {
			return value
		}
		ways := 0
		for root := 1; root <= n; root++ {
			left := rec(root - 1)
			right := rec(n - root)
			ways += left * right
		}
		mem[n] = ways
		return ways
	}
	return rec(n)
}
