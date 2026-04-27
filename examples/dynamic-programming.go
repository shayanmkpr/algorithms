//go:build ignore

package examples

/*
================================================================
Algorithm: Dynamic Programming
================================================================

[EASY] LeetCode 70 - Climbing Stairs
You can climb 1 or 2 steps at a time. How many distinct ways to reach the
top of n stairs?
Example: n = 3  ->  3 (1+1+1, 1+2, 2+1)
Hint: Fibonacci-like recurrence: dp[i] = dp[i-1] + dp[i-2]. Use two rolling
variables for O(1) space.

[EASY] LeetCode 746 - Min Cost Climbing Stairs
cost[i] is the cost of step i. You can start at step 0 or 1; from any
step, take 1 or 2 steps. Return the min cost to reach the top.
Example: cost = [10,15,20]  ->  15
Hint: dp[i] = cost[i] + min(dp[i-1], dp[i-2]). Answer is min(dp[n-1],
dp[n-2]). Use two rolling variables.

[EASY] LeetCode 198 - House Robber
You can't rob two adjacent houses. Maximize total robbed.
Example: nums = [2,7,9,3,1]  ->  12 (rob houses 0, 2, 4)
Hint: dp[i] = max(dp[i-1], dp[i-2] + nums[i]). Two rolling variables: prev
and cur — O(1) space.

[MEDIUM] LeetCode 322 - Coin Change
Given coins of various denominations and amount, return the fewest coins
needed to make up that amount, or -1.
Example: coins=[1,2,5], amount=11  ->  3 (5+5+1)
Hint: dp[a] = min over coins c <= a of dp[a-c] + 1. Initialize dp[0]=0
and dp[a]=amount+1 (a sentinel "infinity").

[HARD] LeetCode 72 - Edit Distance
Given two strings word1, word2, return the minimum operations (insert,
delete, replace) to transform word1 into word2.
Example: word1="horse", word2="ros"  ->  3
Hint: dp[i][j] = edit distance for word1[:i] and word2[:j]. If chars match,
dp[i][j]=dp[i-1][j-1]; else 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]).
*/

// ===================== Answers (Optimal Solutions) =====================

// 1) [EASY] LC 70: Climbing Stairs -- O(n) time, O(1) space
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// 2) [EASY] LC 746: Min Cost Climbing Stairs -- O(n) time, O(1) space
func minCostClimbingStairs(cost []int) int {
	a, b := 0, 0
	for i := 2; i <= len(cost); i++ {
		x := cost[i-1] + b
		y := cost[i-2] + a
		if x < y {
			a, b = b, x
		} else {
			a, b = b, y
		}
	}
	return b
}

// 3) [EASY] LC 198: House Robber -- O(n) time, O(1) space
func rob(nums []int) int {
	prev, cur := 0, 0
	for _, v := range nums {
		take := prev + v
		if take > cur {
			prev, cur = cur, take
		} else {
			prev, cur = cur, cur
		}
	}
	return cur
}

// 4) [MEDIUM] LC 322: Coin Change -- O(amount * len(coins))
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			if c <= a && dp[a-c]+1 < dp[a] {
				dp[a] = dp[a-c] + 1
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// 5) [HARD] LC 72: Edit Distance -- O(m*n) time, O(n) space
func minDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)
	prev := make([]int, n+1)
	for j := 0; j <= n; j++ {
		prev[j] = j
	}
	cur := make([]int, n+1)
	for i := 1; i <= m; i++ {
		cur[0] = i
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				cur[j] = prev[j-1]
			} else {
				m1 := prev[j]
				if cur[j-1] < m1 {
					m1 = cur[j-1]
				}
				if prev[j-1] < m1 {
					m1 = prev[j-1]
				}
				cur[j] = m1 + 1
			}
		}
		prev, cur = cur, prev
	}
	return prev[n]
}
