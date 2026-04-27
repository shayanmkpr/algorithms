//go:build ignore

package examples

/*
================================================================
Algorithm: Greedy Algorithms
================================================================

[EASY] LeetCode 121 - Best Time to Buy and Sell Stock
Given prices[i] = price on day i, return the max profit from one buy + one
sell. If no profit, return 0.
Example: prices = [7,1,5,3,6,4]  ->  5 (buy at 1, sell at 6)
Hint: Track the running min price. For each day, update best profit =
price - minSoFar.

[EASY] LeetCode 860 - Lemonade Change
Customers pay $5, $10, or $20. Each lemonade costs $5. Return true if you
can give the right change to every customer in order, starting with $0.
Example: bills = [5,5,5,10,20]  ->  true
Hint: Track count of $5 and $10 bills. For $20, prefer to use one $10 +
one $5 (greedy: keep $5s for more flexibility); else use three $5s.

[EASY] LeetCode 1221 - Split a String in Balanced Strings
A balanced string has equal counts of 'L' and 'R'. Split s into the maximum
number of balanced substrings.
Example: s = "RLRRLLRLRL"  ->  4
Hint: Greedy counter: +1 for 'R', -1 for 'L'. Each time it returns to 0,
we've closed one balanced piece — increment the counter.

[MEDIUM] LeetCode 55 - Jump Game
Each index has a non-negative jump length. Determine if you can reach the
last index from index 0.
Example: nums = [2,3,1,1,4]  ->  true ; [3,2,1,0,4]  ->  false
Hint: Track the farthest index reachable. If at index i > reachable, return
false. Otherwise update reachable = max(reachable, i + nums[i]).

[HARD] LeetCode 135 - Candy
Children stand in a line with ratings[i]. Each gets at least 1 candy and a
child with a higher rating than a neighbor must get more candies. Return
the minimum total candies.
Example: ratings = [1,0,2]  ->  5 ([2,1,2])
Hint: Two passes. Left-to-right: if ratings[i]>ratings[i-1], candy[i]=
candy[i-1]+1. Right-to-left: if ratings[i]>ratings[i+1], candy[i]=
max(candy[i], candy[i+1]+1). Sum.
*/

// ===================== Answers (Optimal Solutions) =====================

// 1) [EASY] LC 121: Best Time to Buy and Sell Stock -- O(n) time, O(1) space
func maxProfit(prices []int) int {
	const inf = 1 << 30
	minPrice, best := inf, 0
	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else if p-minPrice > best {
			best = p - minPrice
		}
	}
	return best
}

// 2) [EASY] LC 860: Lemonade Change -- O(n) time, O(1) space
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, b := range bills {
		switch b {
		case 5:
			five++
		case 10:
			if five == 0 {
				return false
			}
			five--
			ten++
		default: // 20
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

// 3) [EASY] LC 1221: Split a String in Balanced Strings -- O(n)
func balancedStringSplit(s string) int {
	bal, count := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			bal++
		} else {
			bal--
		}
		if bal == 0 {
			count++
		}
	}
	return count
}

// 4) [MEDIUM] LC 55: Jump Game -- O(n)
func canJump(nums []int) bool {
	reach := 0
	for i, v := range nums {
		if i > reach {
			return false
		}
		if i+v > reach {
			reach = i + v
		}
	}
	return true
}

// 5) [HARD] LC 135: Candy -- O(n) time, O(n) space
func candy(ratings []int) int {
	n := len(ratings)
	c := make([]int, n)
	for i := range c {
		c[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			c[i] = c[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && c[i] <= c[i+1] {
			c[i] = c[i+1] + 1
		}
	}
	total := 0
	for _, v := range c {
		total += v
	}
	return total
}
