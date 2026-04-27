//go:build ignore

package examples

import (
	"container/heap"
	"sort"
)

/*
================================================================
Algorithm: Minimum Spanning Tree (Kruskal / Prim)
================================================================

[EASY] LeetCode 1584 - Min Cost to Connect All Points
Given 2D points, the cost between two points is the Manhattan distance.
Return the minimum cost to connect all points (any two are connected if
there is a path of edges).
Example: points=[[0,0],[2,2],[3,10],[5,2],[7,0]]  ->  20
Hint: Build all pairwise edges with Manhattan distance, sort, and run
Kruskal with Union-Find. Stop after n-1 edges accepted.

[EASY] LeetCode 1102 - Path With Maximum Minimum Value
Find a path from top-left to bottom-right that maximizes the minimum cell
value (the bottleneck).
Example: grid=[[5,4,5],[1,2,6],[7,4,6]]  ->  4
Hint: Sort cells by value descending. Add them one by one (DSU); whenever
top-left and bottom-right become connected, the last added value is the
answer. Same idea as Kruskal's MST.

[EASY] LeetCode 1167 - Minimum Cost to Connect Sticks
Combine sticks two at a time; cost = sum. Return min total cost.
Example: sticks = [2,4,3]  ->  14
Hint: Greedy/MST-flavored: always merge the two smallest. Min-heap; pop
two, push sum, add to total, repeat until one stick remains.

[MEDIUM] LeetCode 1135 - Connecting Cities With Minimum Cost
Given n cities and connections[i] = [a, b, cost], return the minimum cost
to connect all cities, or -1 if impossible.
Example: n=3, connections=[[1,2,5],[1,3,6],[2,3,1]]  ->  6
Hint: Classic Kruskal: sort edges, union endpoints if not yet connected.
After processing, if you've accepted n-1 edges -> total; else -1.

[HARD] LeetCode 1489 - Find Critical and Pseudo-Critical Edges in MST
Return [critical, pseudo] edge indices.
- Critical: removing it strictly increases MST weight.
- Pseudo-critical: appears in some MST but not all.
Example: n=5, edges=[[0,1,1],[1,2,1],[2,3,2],[0,3,2],[0,4,3],[3,4,3],
[1,4,6]]  ->  [[0,1],[2,3,4,5]]
Hint: Compute baseMST. For each edge i: (a) build MST excluding it -> if
weight > base or graph disconnected => critical. Else (b) build MST
forcing i -> if equal to base => pseudo-critical.
*/

// ===================== Answers (Optimal Solutions) =====================

type kdsu struct {
	p, r []int
	cnt  int
}

func newDSU(n int) *kdsu {
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &kdsu{p: p, r: make([]int, n), cnt: n}
}

func (d *kdsu) find(x int) int {
	for d.p[x] != x {
		d.p[x] = d.p[d.p[x]]
		x = d.p[x]
	}
	return x
}

func (d *kdsu) union(a, b int) bool {
	ra, rb := d.find(a), d.find(b)
	if ra == rb {
		return false
	}
	if d.r[ra] < d.r[rb] {
		ra, rb = rb, ra
	}
	d.p[rb] = ra
	if d.r[ra] == d.r[rb] {
		d.r[ra]++
	}
	d.cnt--
	return true
}

// 1) [EASY] LC 1584: Min Cost to Connect All Points -- O(n^2 log n)
func minCostConnectPoints(points [][]int) int {
	n := len(points)
	type edge struct{ u, v, w int }
	edges := make([]edge, 0, n*(n-1)/2)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{i, j, abs(points[i][0]-points[j][0]) + abs(points[i][1]-points[j][1])})
		}
	}
	sort.Slice(edges, func(a, b int) bool { return edges[a].w < edges[b].w })
	d := newDSU(n)
	cost, used := 0, 0
	for _, e := range edges {
		if d.union(e.u, e.v) {
			cost += e.w
			used++
			if used == n-1 {
				break
			}
		}
	}
	return cost
}

// 2) [EASY] LC 1102: Path With Maximum Minimum Value -- O(m*n*log(m*n))
func maximumMinimumPath(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	type cell struct{ v, i, j int }
	cells := make([]cell, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cells = append(cells, cell{grid[i][j], i, j})
		}
	}
	sort.Slice(cells, func(a, b int) bool { return cells[a].v > cells[b].v })

	d := newDSU(m * n)
	added := make([][]bool, m)
	for i := range added {
		added[i] = make([]bool, n)
	}
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	src, dst := 0, m*n-1
	for _, c := range cells {
		added[c.i][c.j] = true
		for _, dr := range dirs {
			ni, nj := c.i+dr[0], c.j+dr[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && added[ni][nj] {
				d.union(c.i*n+c.j, ni*n+nj)
			}
		}
		if d.find(src) == d.find(dst) {
			return c.v
		}
	}
	return 0
}

// 3) [EASY] LC 1167: Minimum Cost to Connect Sticks -- O(n log n)
type minIntHeap []int

func (h minIntHeap) Len() int            { return len(h) }
func (h minIntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minIntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minIntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

func connectSticks(sticks []int) int {
	h := &minIntHeap{}
	*h = append(*h, sticks...)
	heap.Init(h)
	total := 0
	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)
		total += a + b
		heap.Push(h, a+b)
	}
	return total
}

// 4) [MEDIUM] LC 1135: Connecting Cities With Minimum Cost -- O(E log E)
func minimumCost(n int, connections [][]int) int {
	sort.Slice(connections, func(a, b int) bool { return connections[a][2] < connections[b][2] })
	d := newDSU(n + 1) // 1-indexed
	cost, used := 0, 0
	for _, c := range connections {
		if d.union(c[0], c[1]) {
			cost += c[2]
			used++
		}
	}
	if used != n-1 {
		return -1
	}
	return cost
}

// 5) [HARD] LC 1489: Find Critical and Pseudo-Critical Edges in MST -- O(E^2 * α(V))
func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	type idxEdge struct {
		u, v, w, idx int
	}
	es := make([]idxEdge, len(edges))
	for i, e := range edges {
		es[i] = idxEdge{e[0], e[1], e[2], i}
	}
	sorted := make([]idxEdge, len(es))
	copy(sorted, es)
	sort.Slice(sorted, func(a, b int) bool { return sorted[a].w < sorted[b].w })

	mstWeight := func(skip, force int) int {
		d := newDSU(n)
		w := 0
		if force >= 0 {
			e := es[force]
			d.union(e.u, e.v)
			w += e.w
		}
		for _, e := range sorted {
			if e.idx == skip {
				continue
			}
			if d.union(e.u, e.v) {
				w += e.w
			}
		}
		if d.cnt != 1 {
			return 1 << 30
		}
		return w
	}

	base := mstWeight(-1, -1)
	critical, pseudo := []int{}, []int{}
	for i := range es {
		if mstWeight(i, -1) > base {
			critical = append(critical, i)
		} else if mstWeight(-1, i) == base {
			pseudo = append(pseudo, i)
		}
	}
	return [][]int{critical, pseudo}
}
