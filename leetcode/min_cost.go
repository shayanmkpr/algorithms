package main

import "fmt"

func MinCostArr(nums []int) int {
	var max func(i, j int) int
	max = func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}
	mem := make(map[string]int)

	var dfs func(arr []int) int
	dfs = func(arr []int) int {
		if len(arr) < 3 {
			return max(arr[0], arr[1])
		}

		key := fmt.Sprint(arr)
		if cost, ok := mem[key]; ok {
			return cost
		}
		a := append([]int{arr[0]}, arr[2:]...)
		costA := max(arr[1], arr[2]) + dfs(a)

		b := append([]int{arr[1]}, arr[2:]...)
		costB := max(arr[0], arr[2]) + dfs(b)

		c := append([]int{arr[0], arr[1]}, arr[3:]...)
		costC := max(arr[0], arr[1]) + dfs(c)

		res := min(costA, min(costB, costC))
		mem[key] = res
		return res
	}

	return dfs(nums)
}
