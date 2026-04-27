//go:build ignore

package examples

/*
================================================================
Algorithm: Floyd-Warshall (All-Pairs Shortest Paths)
================================================================

[EASY] LeetCode 1334 - Find the City With the Smallest Number of Neighbors
                       at a Threshold Distance
Among n cities, return the city with the fewest reachable cities within
distanceThreshold; ties broken by largest city index.
Example: n=4, edges=[[0,1,3],[1,2,1],[1,3,4],[2,3,1]], threshold=4  ->  3
Hint: Run Floyd-Warshall to get dist[i][j] for all pairs. For each city
count how many j!=i have dist[i][j] <= threshold. Pick min count, larger
index on tie.

[EASY] LeetCode 1971 - Find if Path Exists in Graph (via FW transitive closure)
Return true if a path exists between source and destination in an
undirected graph.
Example: n=3, edges=[[0,1],[1,2],[2,0]], src=0, dst=2  ->  true
Hint: Build a boolean matrix reach[i][j]. For each k, set reach[i][j] |=
reach[i][k] && reach[k][j]. Return reach[source][destination].
(Note: BFS/DSU is faster; this is FW practice.)

[EASY] LeetCode 2101 - Detonate the Maximum Bombs
Each bomb has a 3D position+radius. A bomb b detonates b' if dist(b,b')
<= radius(b). Return the max chain detonations starting from one bomb.
Example: bombs=[[2,1,3],[6,1,4]]  ->  2
Hint: Build directed reach matrix (i can detonate j). Apply FW transitive
closure. Answer is max row count of true values.

[MEDIUM] LeetCode 399 - Evaluate Division
Given equations a/b = v and a list of queries c/d, return c/d's value or
-1 if undeterminable.
Example: equations=[["a","b"],["b","c"]], values=[2.0,3.0],
queries=[["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
                  -> [6.0, 0.5, -1.0, 1.0, -1.0]
Hint: Map variables to indices. Build a value matrix using FW: through
intermediate k, val[i][j] = val[i][k] * val[k][j].

[HARD] LeetCode 1462 - Course Schedule IV
Given prerequisites and queries [u,v], return for each query whether u is
a prerequisite (direct or indirect) of v.
Example: numCourses=2, prerequisites=[[1,0]], queries=[[0,1],[1,0]]
                  ->  [false, true]
Hint: Run a transitive closure with FW: reach[i][j] = reach[i][j] OR
(reach[i][k] AND reach[k][j]). Each query is then O(1).
*/

// ===================== Answers (Optimal Solutions) =====================

// 1) [EASY] LC 1334: Find the City... -- O(n^3)
func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	const inf = 1 << 30
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = inf
			}
		}
	}
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		dist[u][v], dist[v][u] = w, w
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	bestCount, bestCity := n+1, -1
	for i := 0; i < n; i++ {
		c := 0
		for j := 0; j < n; j++ {
			if i != j && dist[i][j] <= distanceThreshold {
				c++
			}
		}
		if c <= bestCount {
			bestCount, bestCity = c, i
		}
	}
	return bestCity
}

// 2) [EASY] LC 1971: Find if Path Exists (FW transitive closure) -- O(n^3)
func validPath(n int, edges [][]int, source, destination int) bool {
	reach := make([][]bool, n)
	for i := range reach {
		reach[i] = make([]bool, n)
		reach[i][i] = true
	}
	for _, e := range edges {
		reach[e[0]][e[1]] = true
		reach[e[1]][e[0]] = true
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if !reach[i][k] {
				continue
			}
			for j := 0; j < n; j++ {
				if reach[k][j] {
					reach[i][j] = true
				}
			}
		}
	}
	return reach[source][destination]
}

// 3) [EASY] LC 2101: Detonate the Maximum Bombs -- O(n^3)
func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	reach := make([][]bool, n)
	for i := range reach {
		reach[i] = make([]bool, n)
		reach[i][i] = true
	}
	for i := 0; i < n; i++ {
		xi, yi, ri := bombs[i][0], bombs[i][1], bombs[i][2]
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			dx, dy := bombs[j][0]-xi, bombs[j][1]-yi
			if dx*dx+dy*dy <= ri*ri {
				reach[i][j] = true
			}
		}
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if !reach[i][k] {
				continue
			}
			for j := 0; j < n; j++ {
				if reach[k][j] {
					reach[i][j] = true
				}
			}
		}
	}
	best := 0
	for i := 0; i < n; i++ {
		c := 0
		for j := 0; j < n; j++ {
			if reach[i][j] {
				c++
			}
		}
		if c > best {
			best = c
		}
	}
	return best
}

// 4) [MEDIUM] LC 399: Evaluate Division -- O(V^3 + Q)
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	id := map[string]int{}
	for _, e := range equations {
		for _, s := range e {
			if _, ok := id[s]; !ok {
				id[s] = len(id)
			}
		}
	}
	n := len(id)
	val := make([][]float64, n)
	for i := range val {
		val[i] = make([]float64, n)
		val[i][i] = 1.0
	}
	for i, e := range equations {
		a, b := id[e[0]], id[e[1]]
		val[a][b] = values[i]
		val[b][a] = 1.0 / values[i]
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if val[i][k] == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				if val[i][j] == 0 && val[k][j] != 0 {
					val[i][j] = val[i][k] * val[k][j]
				}
			}
		}
	}
	out := make([]float64, len(queries))
	for i, q := range queries {
		a, ok1 := id[q[0]]
		b, ok2 := id[q[1]]
		if !ok1 || !ok2 || val[a][b] == 0 {
			out[i] = -1.0
		} else {
			out[i] = val[a][b]
		}
	}
	return out
}

// 5) [HARD] LC 1462: Course Schedule IV -- O(N^3 + Q)
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	r := make([][]bool, numCourses)
	for i := range r {
		r[i] = make([]bool, numCourses)
	}
	for _, p := range prerequisites {
		r[p[0]][p[1]] = true
	}
	for k := 0; k < numCourses; k++ {
		for i := 0; i < numCourses; i++ {
			if !r[i][k] {
				continue
			}
			for j := 0; j < numCourses; j++ {
				if r[k][j] {
					r[i][j] = true
				}
			}
		}
	}
	out := make([]bool, len(queries))
	for i, q := range queries {
		out[i] = r[q[0]][q[1]]
	}
	return out
}
