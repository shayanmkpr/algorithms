package main

import "fmt"

type loc struct {
	x, y int
}

func mynumIslands(grid [][]byte) int {
	fmt.Println(grid)
	dirs := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	seen := map[loc]bool{}
	var result int
	/*
		lets say each element is a node, so it has four children at most.
		if the node is 1, and its child of a 0, then thats an island until
		if the child is water, and it wasnt seen before,
	*/
	var bfs func(grid [][]byte, start loc)
	bfs = func(grid [][]byte, start loc) {
		queue := []loc{start}
		seen[start] = true
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			// logic here
			for i := range dirs {
				neighbour := loc{
					x: curr.x + dirs[i][0],
					y: curr.y + dirs[i][1],
				}
				if neighbour.x < 0 || neighbour.x >= len(grid) || neighbour.y < 0 || neighbour.y >= len(grid[0]) {
					continue
				}
				if !seen[neighbour] {
					seen[neighbour] = true
					queue = append(queue, neighbour)
				}
				seen[neighbour] = true
			}
		}
	}
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == '1' && !seen[loc{x: i, y: j}] {
				result++
				bfs(grid, loc{x: i, y: j})
			}
		}
	}
	return result
}
