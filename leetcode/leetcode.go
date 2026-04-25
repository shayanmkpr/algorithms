package main

import (
	"fmt"
	"slices"
	"strconv"
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	islands := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				islands++
				dfs(grid, r, c)
			}
		}
	}

	return islands
}

func dfs(grid [][]byte, r, c int) {
	rows, cols := len(grid), len(grid[0])

	// Check boundary conditions and if it's not a land cell
	if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '0' {
		return
	}

	// Mark current cell as visited
	grid[r][c] = '0'

	// Explore 4-directionally adjacent cells
	dfs(grid, r+1, c)
	dfs(grid, r-1, c)
	dfs(grid, r, c+1)
	dfs(grid, r, c-1)

	fmt.Println(r, c)
}

type grid struct {
	x int
	y int
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ { // check small squares
		gridMap := make(map[byte]bool)
		for cntr := 0; cntr < 9; cntr++ {
			if gridMap[board[3*(i%3)+cntr%3][3*(i/3)+cntr/3]] {
				return false
			}
			if board[3*(i%3)+cntr%3][3*(i/3)+cntr/3] != '.' {
				gridMap[board[3*(i%3)+cntr%3][3*(i/3)+cntr/3]] = true
			}
		}
	}

	for j := 0; j < 9; j++ {
		gridMap := make(map[byte]bool)
		for cntr := 0; cntr < 9; cntr++ {
			if gridMap[board[cntr][j]] {
				return false
			}
			if board[cntr][j] != '.' {
				gridMap[board[cntr][j]] = true
			}
		}
	}

	for k := 0; k < 9; k++ {
		gridMap := make(map[byte]bool)
		for cntr := 0; cntr < 9; cntr++ {
			if gridMap[board[k][cntr]] {
				return false
			}
			if board[k][cntr] != '.' {
				gridMap[board[k][cntr]] = true
			}
		}
	}

	return true
}

func rotate(matrix [][]int) {
	fmt.Println(len(matrix))
	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			matrix[i][j] = matrix[j][len(matrix)-1-i]
		}
	}
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	N := len(matrix) * len(matrix[0])
	seen := make(map[grid]bool)
	dir := 0 // for moving right
	x, y := 0, 0
	var path []int

	for i := 0; i < N; i++ {
		path = append(path, matrix[y][x])
		seen[grid{x: x, y: y}] = true

		if dir == 0 {
			if x+1 < len(matrix[0]) && !seen[grid{x: x + 1, y: y}] {
				x++
			} else {
				dir = 1
				y++
			}
		} else if dir == 1 {
			if y+1 < len(matrix) && !seen[grid{x: x, y: y + 1}] {
				y++
			} else {
				dir = 2
				x--
			}
		} else if dir == 2 {
			if x-1 >= 0 && !seen[grid{x: x - 1, y: y}] {
				x--
			} else {
				dir = 3
				y--
			}
		} else {
			if y-1 >= 0 && !seen[grid{x: x, y: y - 1}] {
				y--
			} else {
				dir = 0
				x++
			}
		}
	}
	return path
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}

	fmt.Println(queue)

	for len(queue) > 0 {
		size := len(queue)
		level := []int{}

		// process one level
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// instead of adding at the end,
		// add the level at the BEGINNING to get bottom-up order
		result = append([][]int{level}, result...)
	}

	return result
}

func canConstruct(ransomNote string, magazine string) bool {
	magMap := make(map[rune]int)
	for _, ref := range magazine {
		magMap[ref] += 1
	}
	for _, char := range ransomNote {
		if magMap[char] > 0 { // then there is the char
			magMap[char] -= 1
		} else if magMap[char] == 0 {
			return false
		}
	}
	return true
}

func lexicalOrder(n int) []int {
	var result []int
	var dfs func(curr int)
	dfs = func(curr int) {
		if curr > n {
			return
		}

		result = append(result, curr)
		for i := 0; i < 10; i++ {
			next := curr*10 + i
			dfs(next)
		}
	}
	for i := 1; i < 10; i++ {
		dfs(i)
	}
	return result
}

func firstUniqChar(s string) int {
	myMap := make(map[byte]int)
	for i := len(s) - 1; i >= 0; i-- {
		myMap[s[i]] += 1
	}

	for i := range len(s) {
		if myMap[s[i]] == 1 {
			return i
		} else if i == len(s)-1 && myMap[s[i]] != 1 {
			return -1
		}
	}
	return -1
}

func numSquares(n int) int {
	squares := []int{}
	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}

	queue := []int{n}
	visited := make([]bool, n+1)
	visited[n] = true

	depth := 0

	for len(queue) > 0 {
		depth++

		for range queue {
			curr := queue[0]
			queue = queue[1:] // removing the current parent from the queue

			for _, sq := range squares {
				next := curr - sq
				if next < 0 {
					break
				}
				if next == 0 {
					return depth
				}
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
	}
	return depth
}

func countIslands(grid [][]int, k int) int {
	type pos struct {
		x int
		y int
	}

	var queue []pos
	visited := make(map[pos]bool, len(grid)*len(grid[0]))

	islands := 0

	dirs := []pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {

			if grid[i][j] == 0 {
				continue
			}

			start := pos{i, j}
			if visited[start] {
				continue
			}

			// start BFS for a new island
			queue = []pos{start}
			visited[start] = true
			sum := 0

			for len(queue) > 0 {
				curr := queue[0]
				queue = queue[1:]

				sum += grid[curr.x][curr.y]

				for _, d := range dirs {
					x := curr.x + d.x
					y := curr.y + d.y

					if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
						continue
					}

					next := pos{x, y}

					if visited[next] {
						continue
					}

					if grid[x][y] == 0 {
						continue
					}

					visited[next] = true
					queue = append(queue, next)
				}
			}

			if sum%k == 0 {
				fmt.Println(sum)
				islands++
			}
		}
	}

	return islands
}

func sumOfLeftLeaves(root *TreeNode) int {
	result := 0
	start := root
	visited := make(map[*TreeNode]bool)
	queue := []*TreeNode{start}
	visited[start] = true
	for len(queue) > 0 { // while the queue is live
		node := queue[0]
		queue = queue[1:]
		visited[node] = true
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			result += node.Left.Val
		}

		if node.Left != nil && !visited[node.Left] {
			queue = append(queue, node.Left)
		}
		if node.Right != nil && !visited[node.Right] {
			queue = append(queue, node.Right)
		}
	}
	return result
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int

	var dfs func(curr *TreeNode, tempArr []int, tempVal int)
	dfs = func(curr *TreeNode, tempArr []int, tempVal int) {
		tempVal += curr.Val
		tempArr = append(tempArr, curr.Val)
		if curr.Left == nil && curr.Right == nil {
			if tempVal == targetSum {
				goodPath := make([]int, len(tempArr))
				copy(goodPath, tempArr)
				result = append(result, goodPath)
				fmt.Println(tempArr)
				result = append(result, tempArr)
				fmt.Println(tempArr)
			}
			return
		}
		if curr.Left != nil {
			dfs(curr.Left, tempArr, tempVal)
		}
		if curr.Right != nil {
			dfs(curr.Right, tempArr, tempVal)
		}
	}

	dfs(root, []int{}, 0)
	return result
}

func minCost(s string, cost []int) int64 {
	myMap := make(map[rune]int)
	var myMax int
	var totalCost int

	for i, char := range s {
		myMap[char] += cost[i]
	}

	for _, realCost := range myMap {
		totalCost += realCost
		if realCost > myMax {
			myMax = realCost
		}
	}

	return int64(totalCost - myMax)
}

func findDifferentBinaryStringDFS(nums []string) string {
	if len(nums) == 0 {
		return ""
	}
	// given a list of strings that are showing a binary number, give out a string of a binary tree that is not in the given list.

	size := len(nums[0])
	var result string
	seen := make(map[string]bool)
	inputSeen := make(map[string]bool)

	for _, num := range nums {
		inputSeen[num] = true
	}

	var dfs func(curr string)
	dfs = func(curr string) {
		fmt.Println(curr)
		if seen[curr] {
			return
		}

		seen[curr] = true

		if len(curr) == size {
			if !inputSeen[curr] {
				result = curr
			}
			return
		}

		dfs(curr + "1")
		dfs(curr + "0")
	}
	dfs("")
	return result
}

func findDifferentBinaryString(nums []string) string {
	if len(nums) == 0 {
		return ""
	}

	size := len(nums[0])
	seen := make(map[int64]bool)

	// build "111...1"
	maxString := ""
	for i := 0; i < size; i++ {
		maxString += "1"
	}

	// parse inputs
	for _, num := range nums {
		number, err := strconv.ParseInt(num, 2, 64)
		if err != nil {
			continue
		}
		seen[number] = true
	}

	maxVal, _ := strconv.ParseInt(maxString, 2, 64)

	for i := int64(0); i <= maxVal; i++ {
		if !seen[i] {
			bin := strconv.FormatInt(i, 2)
			// pad with leading zeros
			for len(bin) < size {
				bin = "0" + bin
			}
			return bin
		}
	}

	return ""
}

func findBottomLeftValue(root *TreeNode) int {
	maxDepth := -1
	result := 0

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		// first time reaching this depth (leftmost wins)
		if depth > maxDepth {
			maxDepth = depth
			result = node.Val
		}

		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	dfs(root, 0)
	return result
}

func longestContinuousSubstring(s string) int {
	// abcalksjdf
	var i, longest int = 0, 1
	for i < len(s) {
		j := i + 1
		temp := 1
		for j < len(s) && s[j] == s[i]+byte(j-i) {
			temp++
			j++
			if temp > longest {
				longest = temp
			}
		}
		i++
	}
	return longest
}

func minimumSum(n int, k int) int {
	if n == 1 {
		return 1
	}
	fmt.Println(n)
	fmt.Println(k)
	var arr []int
	var kPos, low, high, result int
	for i := range n {
		arr = append(arr, i+1)
		if i+1 == k {
			kPos = i
		}
	}

	if arr[n-1] == k {
		arr = append(arr, arr[n-1]+1)
	}

	high = kPos - 1
	for low < high && high < kPos {
		fmt.Println("low", low)
		fmt.Println("high", high)
		if arr[low]+arr[high] == k {
			arr = slices.Delete(arr, high, high+1)
			arr = append(arr, (arr[len(arr)-1] + 1))
			low++
			high -= 2
		}
		low++
		high--
	}

	for _, val := range arr {
		result += val
	}
	fmt.Println(arr)
	return result
}

func subsetsWithDup(nums []int) [][]int {
	fmt.Println(nums)
	// given an array, give out all the subsets
	var result [][]int
	seen := make(map[int]bool)

	var dfs func(curr []int, length int)
	dfs = func(curr []int, length int) {
		if len(curr) == length {
			// copy slice before saving
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			result = append(result, tmp)
			return
		}

		for i := range nums {
			if seen[i] {
				continue
			}

			seen[i] = true
			dfs(append(curr, nums[i]), length)
			seen[i] = false // backtrack
		}
	}

	dfs([]int{}, 2)

	// for length := range len(nums) {
	// 	dfs([]int{}, length)
	// }

	return result
}

func minStepsDFS(n int) int {
	var minActs int = n
	var dfs func(length int, acts int, clip int, copy bool)
	dfs = func(length int, acts int, clip int, allowCopy bool) {
		if length > n {
			return
		}
		if length == n {
			if acts < minActs {
				minActs = acts
			}
		}
		if length < n {
			// just copy
			if allowCopy {
				dfs(length, acts+1, length, false)
			}
			// paste all
			if clip > 0 {
				dfs(length+clip, acts+1, clip, true)
			}
		}
	}
	dfs(1, 0, 0, true)
	return minActs
}

func addMinimum(word string) int {
	// according to word, we will have to add a, b, or c to make everything a repetition of abc.

	var expect func(c byte) byte
	expect = func(c byte) byte {
		switch c {
		case 'a':
			return 'b'
		case 'b':
			return 'c'
		}
		return 'a'
	}

	fmt.Println(word)
	var result int
	var expected byte = 'a'
	for _, c := range word {
		for c != rune(expected) {
			result++
			expected = expect(expected)
		}
		expected = expect(expected)
	}
	for expected != 'a' {
		result++
		expected = expect(expected)
	}
	return result
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	evenHead := head.Next
	even := evenHead

	for curr != nil && even != nil && even.Next != nil {
		curr.Next = even.Next
		curr = curr.Next
		// now go one step forward
		even.Next = even.Next.Next
		even = even.Next
	}
	curr.Next = evenHead
	return head
}

func twoEditWords(queries []string, dictionary []string) []string {
	if len(queries) == 0 || len(dictionary) == 0 {
		return []string{}
	}
	added := make(map[int]bool)
	var result []string
	for queryIndex, query := range queries {
		for _, word := range dictionary {
			if len(query) != len(word) {
				continue
			}
			cntr := 0
			for i := range query {
				if query[i] != word[i] {
					cntr++
				}
				if cntr > 2 {
					break
				}
				if i == len(query)-1 && added[queryIndex] == false {
					result = append(result, query)
					added[queryIndex] = true
				}
			}
		}
	}
	return result
}
