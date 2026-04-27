//go:build ignore

package examples

import "container/heap"

/*
================================================================
Algorithm: Dijkstra's Algorithm
================================================================

[EASY] LeetCode 743 - Network Delay Time
Given times[i] = [u, v, w] (directed weighted edges) and a starting node k,
return the minimum time for a signal to reach all n nodes (or -1).
Example: times=[[2,1,1],[2,3,1],[3,4,1]], n=4, k=2  ->  2
Hint: Standard Dijkstra. Answer is the max distance among all nodes; if
any node is still infinity, return -1.

[EASY] LeetCode 1514 - Path with Maximum Probability
Undirected graph with edge probabilities. Return the max-probability path
from start to end (0 if none).
Example: edges=[[0,1],[1,2],[0,2]], probs=[0.5,0.5,0.2], start=0, end=2 -> 0.25
Hint: Dijkstra with a MAX-heap on probability. dist[v] = best prob; relax
via dist[u] * weight. Stop when popping `end`.

[EASY] LeetCode 1976 - Number of Ways to Arrive at Destination
Find the number of shortest-path routes from 0 to n-1 in a weighted graph.
Return the count modulo 1e9+7.
Example: roads=[[0,6,7],[0,1,2],[1,2,3],[1,3,3],...]  ->  4
Hint: Run Dijkstra. Maintain ways[v]: when relaxing finds equal distance,
ways[v] += ways[u]; when strictly shorter, ways[v] = ways[u].

[MEDIUM] LeetCode 1631 - Path With Minimum Effort
In an m x n grid of heights, find a path from top-left to bottom-right
minimizing the maximum absolute difference between adjacent cells on the
path.
Example: heights = [[1,2,2],[3,8,2],[5,3,5]]  ->  2
Hint: Treat the grid as a graph where edge cost = |h(a)-h(b)|. Dijkstra,
but the path "cost" is max edge so far, not sum.

[HARD] LeetCode 778 - Swim in Rising Water
At time t, you can swim between two adjacent cells if both have elevation
<= t. Find the minimum t to swim from (0,0) to (n-1,n-1).
Example: grid=[[0,2],[1,3]]  ->  3
Hint: Dijkstra where the distance is max(grid[neighbor], dist[cur]).
Answer is dist[bottom-right].
*/

// ===================== Answers (Optimal Solutions) =====================

type pqItem struct {
	dist, node int
}
type minPQ []pqItem

func (h minPQ) Len() int            { return len(h) }
func (h minPQ) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h minPQ) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minPQ) Push(x interface{}) { *h = append(*h, x.(pqItem)) }
func (h *minPQ) Pop() interface{}   { old := *h; n := len(old); v := old[n-1]; *h = old[:n-1]; return v }

// 1) [EASY] LC 743: Network Delay Time -- O((V+E) log V)
func networkDelayTime(times [][]int, n, k int) int {
	const inf = 1 << 30
	adj := make([][][2]int, n+1)
	for _, e := range times {
		adj[e[0]] = append(adj[e[0]], [2]int{e[1], e[2]})
	}
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = inf
	}
	dist[k] = 0
	pq := &minPQ{{0, k}}
	heap.Init(pq)
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(pqItem)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, nb := range adj[cur.node] {
			if cur.dist+nb[1] < dist[nb[0]] {
				dist[nb[0]] = cur.dist + nb[1]
				heap.Push(pq, pqItem{dist[nb[0]], nb[0]})
			}
		}
	}
	ans := 0
	for i := 1; i <= n; i++ {
		if dist[i] == inf {
			return -1
		}
		if dist[i] > ans {
			ans = dist[i]
		}
	}
	return ans
}

// 2) [EASY] LC 1514: Path with Maximum Probability -- O((V+E) log V)
type probItem struct {
	prob float64
	node int
}
type probMaxPQ []probItem

func (h probMaxPQ) Len() int            { return len(h) }
func (h probMaxPQ) Less(i, j int) bool  { return h[i].prob > h[j].prob }
func (h probMaxPQ) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *probMaxPQ) Push(x interface{}) { *h = append(*h, x.(probItem)) }
func (h *probMaxPQ) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

func maxProbability(n int, edges [][]int, succProb []float64, start, end int) float64 {
	adj := make([][][2]float64, n) // [neighbor, prob]
	for i, e := range edges {
		adj[e[0]] = append(adj[e[0]], [2]float64{float64(e[1]), succProb[i]})
		adj[e[1]] = append(adj[e[1]], [2]float64{float64(e[0]), succProb[i]})
	}
	prob := make([]float64, n)
	prob[start] = 1.0
	pq := &probMaxPQ{{1.0, start}}
	heap.Init(pq)
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(probItem)
		if cur.node == end {
			return cur.prob
		}
		if cur.prob < prob[cur.node] {
			continue
		}
		for _, nb := range adj[cur.node] {
			np := cur.prob * nb[1]
			v := int(nb[0])
			if np > prob[v] {
				prob[v] = np
				heap.Push(pq, probItem{np, v})
			}
		}
	}
	return 0
}

// 3) [EASY] LC 1976: Number of Ways to Arrive at Destination -- O((V+E) log V)
func countPaths(n int, roads [][]int) int {
	const mod = 1_000_000_007
	const inf = 1 << 62
	adj := make([][][2]int, n)
	for _, r := range roads {
		adj[r[0]] = append(adj[r[0]], [2]int{r[1], r[2]})
		adj[r[1]] = append(adj[r[1]], [2]int{r[0], r[2]})
	}
	dist := make([]int, n)
	ways := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[0] = 0
	ways[0] = 1
	pq := &minPQ{{0, 0}}
	heap.Init(pq)
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(pqItem)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, nb := range adj[cur.node] {
			nd := cur.dist + nb[1]
			switch {
			case nd < dist[nb[0]]:
				dist[nb[0]] = nd
				ways[nb[0]] = ways[cur.node]
				heap.Push(pq, pqItem{nd, nb[0]})
			case nd == dist[nb[0]]:
				ways[nb[0]] = (ways[nb[0]] + ways[cur.node]) % mod
			}
		}
	}
	return ways[n-1]
}

// 4) [MEDIUM] LC 1631: Path With Minimum Effort -- O(m*n*log(m*n))
func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	const inf = 1 << 30
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = inf
		}
	}
	dist[0][0] = 0
	pq := &minPQ{{0, 0}}
	heap.Init(pq)
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(pqItem)
		x, y := cur.node/n, cur.node%n
		if cur.dist > dist[x][y] {
			continue
		}
		if x == m-1 && y == n-1 {
			return cur.dist
		}
		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			diff := heights[nx][ny] - heights[x][y]
			if diff < 0 {
				diff = -diff
			}
			cost := cur.dist
			if diff > cost {
				cost = diff
			}
			if cost < dist[nx][ny] {
				dist[nx][ny] = cost
				heap.Push(pq, pqItem{cost, nx*n + ny})
			}
		}
	}
	return 0
}

// 5) [HARD] LC 778: Swim in Rising Water -- O(n^2 log n)
func swimInWater(grid [][]int) int {
	n := len(grid)
	const inf = 1 << 30
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = inf
		}
	}
	dist[0][0] = grid[0][0]
	pq := &minPQ{{grid[0][0], 0}}
	heap.Init(pq)
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(pqItem)
		x, y := cur.node/n, cur.node%n
		if x == n-1 && y == n-1 {
			return cur.dist
		}
		if cur.dist > dist[x][y] {
			continue
		}
		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]
			if nx < 0 || nx >= n || ny < 0 || ny >= n {
				continue
			}
			cost := cur.dist
			if grid[nx][ny] > cost {
				cost = grid[nx][ny]
			}
			if cost < dist[nx][ny] {
				dist[nx][ny] = cost
				heap.Push(pq, pqItem{cost, nx*n + ny})
			}
		}
	}
	return -1
}
