//go:build ignore

package examples

/*
================================================================
Algorithm: DFS - Depth-First Search
================================================================

[EASY] LeetCode 104 - Maximum Depth of Binary Tree
Return the maximum depth (number of nodes on the longest root-to-leaf path).
Example: root = [3,9,20,null,null,15,7]  ->  3
Hint: Recurse: depth(node) = 1 + max(depth(left), depth(right)). Base case:
nil node returns 0.

[EASY] LeetCode 226 - Invert Binary Tree
Invert (mirror) a binary tree.
Example: root = [4,2,7,1,3,6,9]  ->  [4,7,2,9,6,3,1]
Hint: DFS post-order: invert left, invert right, then swap them. Or do it
top-down: swap children first, then recurse. Both are O(n).

[EASY] LeetCode 543 - Diameter of Binary Tree
Return the length (in edges) of the longest path between any two nodes.
Example: root = [1,2,3,4,5]  ->  3 (path 4->2->1->3 or 5->2->1->3)
Hint: DFS returning the height of each subtree. At each node update best =
max(best, leftH + rightH). Return max(leftH, rightH) + 1 to the parent.

[MEDIUM] LeetCode 695 - Max Area of Island
Given a binary grid, return the area of the largest island (4-directional
connectivity).
Example: returns the size of the largest connected component of 1's.
Hint: DFS from each unvisited 1, sum 1 + area in 4 directions, then mark
visited (set to 0) so each cell is counted once.

[HARD] LeetCode 124 - Binary Tree Maximum Path Sum
A path is any sequence of nodes connected by edges, not necessarily through
the root. Return the maximum path sum.
Example: root = [-10,9,20,null,null,15,7]  ->  42 (15+20+7)
Hint: For each node, compute the best "downward gain" = node.Val +
max(0, gainLeft) + max(0, gainRight) is the candidate full path; return
node.Val + max(0, max(left, right)) to the parent (only one branch).
*/

// ===================== Answers (Optimal Solutions) =====================

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 1) [EASY] LC 104: Maximum Depth of Binary Tree -- O(n)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}

// 2) [EASY] LC 226: Invert Binary Tree -- O(n)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// 3) [EASY] LC 543: Diameter of Binary Tree -- O(n)
func diameterOfBinaryTree(root *TreeNode) int {
	best := 0
	var depth func(n *TreeNode) int
	depth = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l := depth(n.Left)
		r := depth(n.Right)
		if l+r > best {
			best = l + r
		}
		if l > r {
			return l + 1
		}
		return r + 1
	}
	depth(root)
	return best
}

// 4) [MEDIUM] LC 695: Max Area of Island -- O(m*n)
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return 1 + dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)
	}
	best := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if a := dfs(i, j); a > best {
					best = a
				}
			}
		}
	}
	return best
}

// 5) [HARD] LC 124: Binary Tree Maximum Path Sum -- O(n)
func maxPathSum(root *TreeNode) int {
	const negInf = -1 << 31
	best := negInf
	var gain func(n *TreeNode) int
	gain = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		l := gain(n.Left)
		if l < 0 {
			l = 0
		}
		r := gain(n.Right)
		if r < 0 {
			r = 0
		}
		if n.Val+l+r > best {
			best = n.Val + l + r
		}
		if l > r {
			return n.Val + l
		}
		return n.Val + r
	}
	gain(root)
	return best
}
