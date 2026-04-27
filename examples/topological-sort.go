//go:build ignore

package examples

/*
================================================================
Algorithm: Topological Sort (Kahn's BFS)
================================================================

[EASY] LeetCode 207 - Course Schedule
Given numCourses and prerequisites [a,b] (b -> a), return true if you can
finish all courses (i.e., the graph is a DAG).
Example: 2, [[1,0]]  ->  true ; 2, [[1,0],[0,1]]  ->  false
Hint: Build adjacency + indegree. Push indegree-0 into a queue. Pop, count,
and decrement neighbors' indegree. If processed count == numCourses,
no cycle.

[EASY] LeetCode 1557 - Minimum Number of Vertices to Reach All Nodes
Given a DAG, return the smallest set of vertices from which all nodes are
reachable.
Example: n=6, edges=[[0,1],[0,2],[2,5],[3,4],[4,2]]  ->  [0,3]
Hint: A node is required iff it has indegree 0 (nothing else can reach it).
Just count indegrees and collect nodes with indegree == 0.

[EASY] LeetCode 2115 - Find All Possible Recipes from Given Supplies
Given recipes, ingredient lists, and initial supplies, return all recipes
you can ultimately make.
Example: recipes=["bread"], ing=[["yeast","flour"]], supplies=["yeast","flour"] -> ["bread"]
Hint: Treat each ingredient/recipe as a node. Edge from each ingredient to
recipes that need it. Run Kahn's; supplies start with indegree 0. When a
recipe's indegree drops to 0, it's craftable.

[MEDIUM] LeetCode 210 - Course Schedule II
Same setup, but return any valid course order. If impossible, return [].
Example: 4, [[1,0],[2,0],[3,1],[3,2]]  ->  [0,2,1,3] (one valid answer)
Hint: Same Kahn's algorithm; record nodes in the order they leave the
queue. If the order's length < numCourses, there is a cycle.

[HARD] LeetCode 269 - Alien Dictionary
A list of words is sorted by an unknown alphabet. Return any valid order
of letters used, or "" if impossible.
Example: ["wrt","wrf","er","ett","rftt"]  ->  "wertf"
Hint: For each adjacent word pair, find the first differing char to add an
edge a->b. Watch out for the prefix-conflict case ("abc","ab" -> invalid).
Then run Kahn's; if a cycle exists, return "".
*/

// ===================== Answers (Optimal Solutions) =====================

// 1) [EASY] LC 207: Course Schedule -- O(V + E)
func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, p := range prerequisites {
		adj[p[1]] = append(adj[p[1]], p[0])
		indeg[p[0]]++
	}
	q := []int{}
	for i, d := range indeg {
		if d == 0 {
			q = append(q, i)
		}
	}
	done := 0
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		done++
		for _, y := range adj[x] {
			indeg[y]--
			if indeg[y] == 0 {
				q = append(q, y)
			}
		}
	}
	return done == numCourses
}

// 2) [EASY] LC 1557: Minimum Number of Vertices to Reach All Nodes -- O(V + E)
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	indeg := make([]int, n)
	for _, e := range edges {
		indeg[e[1]]++
	}
	res := []int{}
	for i, d := range indeg {
		if d == 0 {
			res = append(res, i)
		}
	}
	return res
}

// 3) [EASY] LC 2115: Find All Possible Recipes from Given Supplies -- O(V + E)
func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	indeg := make(map[string]int)
	adj := make(map[string][]string)
	isRecipe := make(map[string]bool)
	for i, r := range recipes {
		isRecipe[r] = true
		indeg[r] = len(ingredients[i])
		for _, ing := range ingredients[i] {
			adj[ing] = append(adj[ing], r)
		}
	}
	q := append([]string{}, supplies...)
	out := []string{}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, nb := range adj[cur] {
			indeg[nb]--
			if indeg[nb] == 0 {
				q = append(q, nb)
				if isRecipe[nb] {
					out = append(out, nb)
				}
			}
		}
	}
	return out
}

// 4) [MEDIUM] LC 210: Course Schedule II -- O(V + E)
func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, p := range prerequisites {
		adj[p[1]] = append(adj[p[1]], p[0])
		indeg[p[0]]++
	}
	q := []int{}
	for i, d := range indeg {
		if d == 0 {
			q = append(q, i)
		}
	}
	order := make([]int, 0, numCourses)
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		order = append(order, x)
		for _, y := range adj[x] {
			indeg[y]--
			if indeg[y] == 0 {
				q = append(q, y)
			}
		}
	}
	if len(order) != numCourses {
		return []int{}
	}
	return order
}

// 5) [HARD] LC 269: Alien Dictionary -- O(C) where C is total chars
func alienOrder(words []string) string {
	indeg := make(map[byte]int)
	adj := make(map[byte]map[byte]bool)
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			if _, ok := indeg[w[i]]; !ok {
				indeg[w[i]] = 0
				adj[w[i]] = map[byte]bool{}
			}
		}
	}
	for i := 0; i+1 < len(words); i++ {
		a, b := words[i], words[i+1]
		k := len(a)
		if len(b) < k {
			k = len(b)
		}
		mismatch := false
		for j := 0; j < k; j++ {
			if a[j] != b[j] {
				if !adj[a[j]][b[j]] {
					adj[a[j]][b[j]] = true
					indeg[b[j]]++
				}
				mismatch = true
				break
			}
		}
		if !mismatch && len(a) > len(b) {
			return ""
		}
	}
	q := []byte{}
	for c, d := range indeg {
		if d == 0 {
			q = append(q, c)
		}
	}
	out := []byte{}
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		out = append(out, c)
		for nb := range adj[c] {
			indeg[nb]--
			if indeg[nb] == 0 {
				q = append(q, nb)
			}
		}
	}
	if len(out) != len(indeg) {
		return ""
	}
	return string(out)
}
