package main

func maxAreaOfIsland(grid [][]int) int {
	type loc struct {
		x int
		y int
	}
	var result, local int
	dirs := []loc{
		{-1, 0}, {1, 0}, {0, 1}, {0, -1},
	}
	seen := map[loc]bool{}
	var bfs func(grid [][]int, start loc) int
	bfs = func(grid [][]int, start loc) int {
		local = 0
		queue := []loc{start}
		seen[start] = true
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			// logic on curr
			local++
			for _, dir := range dirs {
				neighbour := loc{x: curr.x + dir.x, y: curr.y + dir.y}
				if neighbour.x >= len(grid[0]) || neighbour.x < 0 || neighbour.y >= len(grid) || neighbour.y < 0 {
					continue
				}
				if !seen[neighbour] && grid[neighbour.y][neighbour.x] != 0 {
					queue = append(queue, neighbour)
					seen[neighbour] = true
				}
			}
		}
		return local
		// start with 1s only
	}
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == 1 && !seen[loc{x: i, y: j}] {
				localMax := bfs(grid, loc{x: i, y: j})
				if localMax > result {
					result = localMax
				}
			}
		}
	}
	return result
}
