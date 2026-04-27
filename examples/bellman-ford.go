//go:build ignore

package examples

/*
================================================================
Algorithm: Bellman-Ford
================================================================

[EASY] LeetCode 787 - Cheapest Flights Within K Stops
Given n cities and flights[i] = [from, to, price], find the cheapest price
from src to dst with at most k stops. Return -1 if impossible.
Example: n=4, flights=[[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]],
src=0, dst=3, k=1  ->  700
Hint: Run Bellman-Ford for exactly k+1 iterations. Snapshot distances each
round so updates only use last round's values (limits "edge count").

[EASY] LeetCode 743 - Network Delay Time (via Bellman-Ford)
Same as the Dijkstra version: directed weighted edges, return time for the
signal from k to reach all nodes (or -1).
Example: times=[[2,1,1],[2,3,1],[3,4,1]], n=4, k=2  ->  2
Hint: Initialize dist[k]=0, others = inf. Relax every edge n-1 times.
Answer = max(dist) if all finite, else -1. Demonstrates BF on
non-negative weights.

[EASY] LeetCode 1334 - Find the City... Threshold Distance (via BF)
Among n cities, return the city with the fewest neighbors reachable within
distanceThreshold (ties: largest index).
Example: n=4, edges=[[0,1,3],[1,2,1],[1,3,4],[2,3,1]], threshold=4  ->  3
Hint: For each source, run Bellman-Ford to get all distances, then count
how many are <= threshold. Pick the city with the smallest count, larger
index on tie.

[MEDIUM] LeetCode 1514 - Path with Maximum Probability
Undirected graph with edge probabilities. Return the max probability path
from start to end (0 if none).
Example: n=3, edges=[[0,1],[1,2],[0,2]], succProb=[0.5,0.5,0.2], start=0,
end=2  ->  0.25
Hint: Bellman-Ford-like relaxation: for each edge (u,v,p), if prob[u]*p
> prob[v], update. Repeat until no updates (or n-1 times). Multiplication
of probabilities replaces sum of weights.

[HARD] LeetCode 1928 - Minimum Cost to Reach Destination in Time
Given an undirected weighted graph and passingFees[i] for each city, find
the min total fees from city 0 to city n-1 within maxTime, or -1.
Example: maxTime=30, edges=[[0,1,10],[1,2,10],[2,5,10],[0,3,1],[3,4,10],
[4,5,15]], fees=[5,1,2,20,20,3]  ->  11
Hint: 2D Bellman-Ford / DP: dp[t][v] = min fees reaching v in <= t time.
Iterate t=1..maxTime; for each edge relax dp[t][v] using dp[t-w][u]+fee[v].
*/

// ===================== Answers (Optimal Solutions) =====================

// 1) [EASY] LC 787: Cheapest Flights Within K Stops -- O((k+1)*E)
func findCheapestPrice(n int, flights [][]int, src, dst, k int) int {
	const inf = 1 << 30
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[src] = 0
	for i := 0; i <= k; i++ {
		next := make([]int, n)
		copy(next, dist)
		for _, f := range flights {
			u, v, w := f[0], f[1], f[2]
			if dist[u] != inf && dist[u]+w < next[v] {
				next[v] = dist[u] + w
			}
		}
		dist = next
	}
	if dist[dst] == inf {
		return -1
	}
	return dist[dst]
}

// 2) [EASY] LC 743: Network Delay Time (Bellman-Ford) -- O(V*E)
func networkDelayTime(times [][]int, n, k int) int {
	const inf = 1 << 30
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = inf
	}
	dist[k] = 0
	for i := 0; i < n-1; i++ {
		updated := false
		for _, e := range times {
			u, v, w := e[0], e[1], e[2]
			if dist[u] != inf && dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				updated = true
			}
		}
		if !updated {
			break
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

// 3) [EASY] LC 1334: Find the City (Bellman-Ford per source) -- O(V^2 * E)
func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	const inf = 1 << 30
	bf := func(src int) []int {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = inf
		}
		dist[src] = 0
		for i := 0; i < n-1; i++ {
			updated := false
			for _, e := range edges {
				u, v, w := e[0], e[1], e[2]
				if dist[u] != inf && dist[u]+w < dist[v] {
					dist[v] = dist[u] + w
					updated = true
				}
				if dist[v] != inf && dist[v]+w < dist[u] {
					dist[u] = dist[v] + w
					updated = true
				}
			}
			if !updated {
				break
			}
		}
		return dist
	}
	bestCount, bestCity := n+1, -1
	for i := 0; i < n; i++ {
		d := bf(i)
		c := 0
		for j, dv := range d {
			if j != i && dv <= distanceThreshold {
				c++
			}
		}
		if c <= bestCount {
			bestCount, bestCity = c, i
		}
	}
	return bestCity
}

// 4) [MEDIUM] LC 1514: Path with Maximum Probability -- O(V*E) worst case
func maxProbability(n int, edges [][]int, succProb []float64, start, end int) float64 {
	prob := make([]float64, n)
	prob[start] = 1.0
	for iter := 0; iter < n-1; iter++ {
		updated := false
		for i, e := range edges {
			u, v := e[0], e[1]
			p := succProb[i]
			if prob[u]*p > prob[v] {
				prob[v] = prob[u] * p
				updated = true
			}
			if prob[v]*p > prob[u] {
				prob[u] = prob[v] * p
				updated = true
			}
		}
		if !updated {
			break
		}
	}
	return prob[end]
}

// 5) [HARD] LC 1928: Minimum Cost to Reach Destination in Time -- O(maxTime * E)
func minCost(maxTime int, edges [][]int, passingFees []int) int {
	n := len(passingFees)
	const inf = 1 << 30
	dp := make([][]int, maxTime+1)
	for t := 0; t <= maxTime; t++ {
		dp[t] = make([]int, n)
		for v := 0; v < n; v++ {
			dp[t][v] = inf
		}
	}
	dp[0][0] = passingFees[0]
	for t := 1; t <= maxTime; t++ {
		for _, e := range edges {
			u, v, w := e[0], e[1], e[2]
			if w > t {
				continue
			}
			if dp[t-w][u]+passingFees[v] < dp[t][v] {
				dp[t][v] = dp[t-w][u] + passingFees[v]
			}
			if dp[t-w][v]+passingFees[u] < dp[t][u] {
				dp[t][u] = dp[t-w][v] + passingFees[u]
			}
		}
	}
	best := inf
	for t := 0; t <= maxTime; t++ {
		if dp[t][n-1] < best {
			best = dp[t][n-1]
		}
	}
	if best == inf {
		return -1
	}
	return best
}
