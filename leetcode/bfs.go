package main

type node struct {
	key   string
	value int
}

// func bfs(graph map[node][]node, start node) map[string]int {
// 	result := map[string]int{}
// 	seen := map[node]bool{}
// 	queue := []node{start}
// 	curr := start
// 	seen[start] = true
// 	for len(queue) > 0 {
// 		curr = queue[0]
// 		queue = queue[1:]
// 		result[curr.key] = curr.value
// 		for _, neighbour := range graph[curr] {
// 			if !seen[neighbour] {
// 				seen[neighbour] = true
// 				queue = append(queue, neighbour)
// 			}
// 		}
// 	}
// 	return result
// }

func bfs(grid map[node][]node, start node) {
	seen := map[node]bool{}
	seen[start] = true
	queue := []node{start}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		seen[curr] = true
		// logic with curr
		for _, neighbour := range grid[curr] {
			// handle the neighbours
			if !seen[neighbour] {
				queue = append(queue, neighbour)
				seen[neighbour] = true
				// logic with neighbour
			}
		}
	}
}
