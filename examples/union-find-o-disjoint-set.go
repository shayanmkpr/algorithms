//go:build ignore

package examples

/*
================================================================
Algorithm: Union-Find / Disjoint Set Union (DSU)
================================================================

[EASY] LeetCode 547 - Number of Provinces
You have an n x n adjacency matrix isConnected[i][j] = 1 if cities i and j
are directly connected. Return the number of connected components.
Example: isConnected = [[1,1,0],[1,1,0],[0,0,1]]  ->  2
Hint: Initialize parent[i]=i. For each i<j with isConnected[i][j]=1, union.
Count distinct roots at the end.

[EASY] LeetCode 1971 - Find if Path Exists in Graph
Given an undirected graph, return true if a path exists from source to
destination.
Example: n=3, edges=[[0,1],[1,2],[2,0]], src=0, dst=2  ->  true
Hint: Union all edges. Path exists iff find(source) == find(destination).

[EASY] LeetCode 990 - Satisfiability of Equality Equations
Given equations like "a==b" and "b!=c", decide if they can all hold.
Example: ["a==b","b!=a"]  ->  false
Hint: Process all "==" first to merge groups. Then check every "!=":
if find(x) == find(y), it's a contradiction.

[MEDIUM] LeetCode 684 - Redundant Connection
Given an undirected graph that started as a tree but had one extra edge
added, return the edge that should be removed (the one creating the cycle).
Example: edges = [[1,2],[1,3],[2,3]]  ->  [2,3]
Hint: Iterate edges. Use union-find: if find(u)==find(v), this edge closes
a cycle and is the answer. Otherwise union(u, v).

[HARD] LeetCode 685 - Redundant Connection II (directed version)
Same as above but the graph is directed. Return the edge that, if removed,
results in a rooted tree.
Example: edges = [[1,2],[1,3],[2,3]]  ->  [2,3]
Hint: There are two failure cases: (a) a node has two parents, (b) a cycle.
Detect a node with two incoming edges; mark candidates A and B. Skip A and
run DSU. If valid tree -> answer is A; else answer is B (or the cycle edge).
*/

// ===================== Answers (Optimal Solutions) =====================

type DSU struct {
	parent, rank []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &DSU{parent: p, rank: make([]int, n)}
}

func (d *DSU) Find(x int) int {
	for d.parent[x] != x {
		d.parent[x] = d.parent[d.parent[x]] // path compression
		x = d.parent[x]
	}
	return x
}

func (d *DSU) Union(a, b int) bool {
	ra, rb := d.Find(a), d.Find(b)
	if ra == rb {
		return false
	}
	if d.rank[ra] < d.rank[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	if d.rank[ra] == d.rank[rb] {
		d.rank[ra]++
	}
	return true
}

// 1) [EASY] LC 547: Number of Provinces -- ~O(n^2 * α(n))
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	d := NewDSU(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				d.Union(i, j)
			}
		}
	}
	roots := map[int]struct{}{}
	for i := 0; i < n; i++ {
		roots[d.Find(i)] = struct{}{}
	}
	return len(roots)
}

// 2) [EASY] LC 1971: Find if Path Exists in Graph -- ~O((V+E) * α(V))
func validPath(n int, edges [][]int, source, destination int) bool {
	d := NewDSU(n)
	for _, e := range edges {
		d.Union(e[0], e[1])
	}
	return d.Find(source) == d.Find(destination)
}

// 3) [EASY] LC 990: Satisfiability of Equality Equations -- ~O(N * α)
func equationsPossible(equations []string) bool {
	d := NewDSU(26)
	for _, eq := range equations {
		if eq[1] == '=' {
			d.Union(int(eq[0]-'a'), int(eq[3]-'a'))
		}
	}
	for _, eq := range equations {
		if eq[1] == '!' && d.Find(int(eq[0]-'a')) == d.Find(int(eq[3]-'a')) {
			return false
		}
	}
	return true
}

// 4) [MEDIUM] LC 684: Redundant Connection -- ~O(N * α(N))
func findRedundantConnection(edges [][]int) []int {
	d := NewDSU(len(edges) + 1)
	for _, e := range edges {
		if !d.Union(e[0], e[1]) {
			return e
		}
	}
	return nil
}

// 5) [HARD] LC 685: Redundant Connection II -- ~O(N * α(N))
func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	var candA, candB []int
	for _, e := range edges {
		u, v := e[0], e[1]
		if parent[v] != 0 {
			candA = []int{parent[v], v} // earlier edge to v
			candB = []int{u, v}         // current edge to v
			e[0], e[1] = 0, 0           // mark candB as ignored
		} else {
			parent[v] = u
		}
	}
	d := NewDSU(n + 1)
	for _, e := range edges {
		if e[0] == 0 && e[1] == 0 {
			continue
		}
		if !d.Union(e[0], e[1]) {
			if candA == nil {
				return e
			}
			return candA
		}
	}
	return candB
}
