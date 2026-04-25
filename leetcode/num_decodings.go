package main

import "strconv"

func numDecodings(s string) int {
	// we want to find out the number of wys that we can decode numbers to strings of words
	// typical back tracking problem, with a touch of confusion.

	var paths [][]string
	var visited map[string]bool

	var backtracking func(numbers string)
	backtracking = func(numbers string) {
		// check if we are at the goal, if we are, then append the path to the paths
		if len(numbers) == 0 {
			paths = append(paths, []string{})
			return
		}

		for i := 1; i <= 2; i++ {
			if i > len(numbers) {
				break
			}
			chosen := numbers[:i]
			if visited[chosen] {
				continue
			}
			visited[chosen] = true
			// now check if the number is in the valid window
			if n, err := strconv.Atoi(chosen); err == nil && n >= 1 && n <= 26 {
				backtracking(numbers[i:])
			}
			delete(visited, chosen)
		}
	}

	return len(paths)
}
