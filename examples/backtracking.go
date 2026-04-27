//go:build ignore

package examples

/*
================================================================
Algorithm: Backtracking
================================================================

[EASY] LeetCode 78 - Subsets
Return all possible subsets (the power set) of distinct integers nums.
Example: nums = [1,2,3]  ->  [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
Hint: For each element, "take it or skip it". Push current subset on entry,
recurse from i+1 onwards (no duplicates that way). Append a copy each time.

[EASY] LeetCode 22 - Generate Parentheses
Given n, generate all combinations of n pairs of well-formed parentheses.
Example: n = 3 -> ["((()))","(()())","(())()","()(())","()()()"]
Hint: Track open and close counts. Add '(' if open < n. Add ')' if close <
open. Stop when length == 2n. Undo (pop) before next branch.

[EASY] LeetCode 17 - Letter Combinations of a Phone Number
Given digits "2"-"9", return all possible letter combinations the number
could represent (phone keypad).
Example: digits = "23"  ->  ["ad","ae","af","bd","be","bf","cd","ce","cf"]
Hint: Map each digit to its letters. Recurse digit by digit; at each level
loop over its letters, append, recurse, pop.

[MEDIUM] LeetCode 46 - Permutations
Given an array of distinct integers, return all possible permutations.
Example: nums = [1,2,3]  ->  6 permutations
Hint: Maintain a `used` boolean array. At each level pick any unused
number, mark it used, recurse, then unmark before trying the next pick.

[HARD] LeetCode 51 - N-Queens
Place n queens on an n x n board so that no two attack each other. Return
all distinct configurations.
Example: n = 4  ->  2 boards.
Hint: Place queens row by row. Track 3 sets: used columns, used (r-c)
diagonals, used (r+c) anti-diagonals. Place, recurse, undo.
*/

// ===================== Answers (Optimal Solutions) =====================

// 1) [EASY] LC 78: Subsets -- O(n * 2^n)
func subsets(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}
	var bt func(i int)
	bt = func(i int) {
		cp := make([]int, len(cur))
		copy(cp, cur)
		res = append(res, cp)
		for j := i; j < len(nums); j++ {
			cur = append(cur, nums[j])
			bt(j + 1)
			cur = cur[:len(cur)-1]
		}
	}
	bt(0)
	return res
}

// 2) [EASY] LC 22: Generate Parentheses -- O(C(n) * n) where C(n) is Catalan
func generateParenthesis(n int) []string {
	res := []string{}
	cur := make([]byte, 0, 2*n)
	var bt func(open, close int)
	bt = func(open, close int) {
		if len(cur) == 2*n {
			res = append(res, string(cur))
			return
		}
		if open < n {
			cur = append(cur, '(')
			bt(open+1, close)
			cur = cur[:len(cur)-1]
		}
		if close < open {
			cur = append(cur, ')')
			bt(open, close+1)
			cur = cur[:len(cur)-1]
		}
	}
	bt(0, 0)
	return res
}

// 3) [EASY] LC 17: Letter Combinations of a Phone Number -- O(4^n * n)
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	pad := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}
	res := []string{}
	cur := make([]byte, 0, len(digits))
	var bt func(i int)
	bt = func(i int) {
		if i == len(digits) {
			res = append(res, string(cur))
			return
		}
		for j := 0; j < len(pad[digits[i]]); j++ {
			cur = append(cur, pad[digits[i]][j])
			bt(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	bt(0)
	return res
}

// 4) [MEDIUM] LC 46: Permutations -- O(n * n!)
func permute(nums []int) [][]int {
	res := [][]int{}
	used := make([]bool, len(nums))
	cur := []int{}
	var bt func()
	bt = func() {
		if len(cur) == len(nums) {
			cp := make([]int, len(cur))
			copy(cp, cur)
			res = append(res, cp)
			return
		}
		for i, v := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			cur = append(cur, v)
			bt()
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}
	bt()
	return res
}

// 5) [HARD] LC 51: N-Queens -- O(n!)
func solveNQueens(n int) [][]string {
	res := [][]string{}
	cols := make([]bool, n)
	d1 := make([]bool, 2*n) // r - c + n
	d2 := make([]bool, 2*n) // r + c
	queens := make([]int, n)

	var bt func(r int)
	bt = func(r int) {
		if r == n {
			board := make([]string, n)
			for i := 0; i < n; i++ {
				row := make([]byte, n)
				for j := range row {
					row[j] = '.'
				}
				row[queens[i]] = 'Q'
				board[i] = string(row)
			}
			res = append(res, board)
			return
		}
		for c := 0; c < n; c++ {
			if cols[c] || d1[r-c+n] || d2[r+c] {
				continue
			}
			cols[c], d1[r-c+n], d2[r+c] = true, true, true
			queens[r] = c
			bt(r + 1)
			cols[c], d1[r-c+n], d2[r+c] = false, false, false
		}
	}
	bt(0)
	return res
}
