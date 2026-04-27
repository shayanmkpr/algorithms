//go:build ignore

package examples

/*
================================================================
Algorithm: BFS - Breadth-First Search
================================================================

[EASY] LeetCode 102 - Binary Tree Level Order Traversal
Return the values of the tree, grouped by level (top to bottom).
Example: root = [3,9,20,null,null,15,7]  ->  [[3],[9,20],[15,7]]
Hint: Use a queue. For each "level", record the current queue size and
process exactly that many nodes before moving to the next level.

[EASY] LeetCode 1971 - Find if Path Exists in Graph
Given an undirected graph and two nodes source/destination, return true if
a path exists.
Example: n=3, edges=[[0,1],[1,2],[2,0]], src=0, dst=2  ->  true
Hint: Build adjacency list. BFS from source; mark visited; return true the
moment you dequeue destination.

[EASY] LeetCode 690 - Employee Importance
Given employees [id, importance, [subordinate ids]], return total importance
(self + all subordinates transitively) of a given id.
Example: [[1,5,[2,3]],[2,3,[]],[3,3,[]]], id=1  ->  11
Hint: Map id -> employee. BFS from the starting id; sum importance values
along the way and enqueue every subordinate.

[MEDIUM] LeetCode 200 - Number of Islands
Given a grid of '1' (land) and '0' (water), return the number of islands.
Example: grid = [["1","1","0"],["1","0","0"],["0","0","1"]]  ->  2
Hint: Scan the grid. When you find a '1', BFS from there and mark all
connected land as visited (or '0'). Increment the island counter once.

[HARD] LeetCode 127 - Word Ladder
Given beginWord, endWord, and wordList, return the length of the shortest
transformation sequence (each step changes one letter and must be in the
list). Return 0 if impossible.
Example: begin="hit", end="cog", list=["hot","dot","dog","lot","log","cog"] -> 5
Hint: Build adjacency lazily via "*" patterns ("h*t", "ho*", "*it"). BFS
from beginWord; the first time you pop endWord, return the level.
*/

// ===================== Answers (Optimal Solutions) =====================

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 1) [EASY] LC 102: Level Order Traversal -- O(n)
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		level := make([]int, 0, size)
		for i := 0; i < size; i++ {
			n := q[i]
			level = append(level, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		q = q[size:]
		res = append(res, level)
	}
	return res
}

// 2) [EASY] LC 1971: Find if Path Exists in Graph -- O(V + E)
func validPath(n int, edges [][]int, source, destination int) bool {
	if source == destination {
		return true
	}
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	visited := make([]bool, n)
	visited[source] = true
	q := []int{source}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, nb := range adj[cur] {
			if nb == destination {
				return true
			}
			if !visited[nb] {
				visited[nb] = true
				q = append(q, nb)
			}
		}
	}
	return false
}

// 3) [EASY] LC 690: Employee Importance -- O(n)
type Employee struct {
	Id, Importance int
	Subordinates   []int
}

func getImportance(employees []*Employee, id int) int {
	idx := make(map[int]*Employee, len(employees))
	for _, e := range employees {
		idx[e.Id] = e
	}
	total := 0
	q := []int{id}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		e := idx[cur]
		if e == nil {
			continue
		}
		total += e.Importance
		q = append(q, e.Subordinates...)
	}
	return total
}

// 4) [MEDIUM] LC 200: Number of Islands -- O(m*n)
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '1' {
				continue
			}
			count++
			q := [][2]int{{i, j}}
			grid[i][j] = '0'
			for len(q) > 0 {
				cell := q[0]
				q = q[1:]
				for _, d := range dirs {
					x, y := cell[0]+d[0], cell[1]+d[1]
					if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == '1' {
						grid[x][y] = '0'
						q = append(q, [2]int{x, y})
					}
				}
			}
		}
	}
	return count
}

// 5) [HARD] LC 127: Word Ladder -- O(N * L^2)
func ladderLength(beginWord, endWord string, wordList []string) int {
	dict := make(map[string]bool, len(wordList))
	for _, w := range wordList {
		dict[w] = true
	}
	if !dict[endWord] {
		return 0
	}
	q := []string{beginWord}
	visited := map[string]bool{beginWord: true}
	steps := 1
	for len(q) > 0 {
		next := []string{}
		for _, w := range q {
			if w == endWord {
				return steps
			}
			b := []byte(w)
			for i := 0; i < len(b); i++ {
				orig := b[i]
				for c := byte('a'); c <= 'z'; c++ {
					if c == orig {
						continue
					}
					b[i] = c
					nw := string(b)
					if dict[nw] && !visited[nw] {
						visited[nw] = true
						next = append(next, nw)
					}
				}
				b[i] = orig
			}
		}
		q = next
		steps++
	}
	return 0
}
