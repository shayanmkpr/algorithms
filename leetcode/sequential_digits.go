package main

func sequentialDigits(low int, high int) []int {
	u := []int{}
	for i := 12; i <= 123456789; i++ {
		u = append(u, i)
	}

	var search func(l, h, target int) int
	search = func(l, h, target int) int {
		i := (l + h) / 2
		if u[i] > target && u[i-1] < target {
			return i
		}
		if u[i] > target && u[i-1] > target {
			return search(l, i, target)
		}
		if u[i] < target && u[i-1] < target {
			return search(i, h, target)
		}
		return i
	}
	lowIdx := search(0, len(u), low)
	highIdx := search(0, len(u), high)

	return u[lowIdx:highIdx]
}
