package main

/*
there are two different approaches.
1. Top-down: Allocating to the memory with forward movement. Could cuse overflow, slow map lookup.
2. Buttom-up: Pre-allocating memroy with a fixed size array. Harder to implement but fater and safer.
*/

func pascalSecond(rowIndex int) []int {
	rows := [][]int{[]int{1}}
	seen := map[int][]int{}
	for i := 1; i <= rowIndex; i++ {
		if _, ok := seen[i]; !ok {
			var row []int
			for j := 0; j <= i; j++ {
				row = append(row, 0)
				if j == 0 || j == i {
					row[j] = 1
				}
				if j > 0 && j < i {
					row[j] = rows[i-1][j-1] + rows[i-1][j]
				} else if j == 0 {
					row[j] = rows[i-1][j]
				} else if j == i {
					row[j] = rows[i-1][j-1]
				}
			}
			seen[i] = row
		}
		rows = append(rows, seen[i])
	}
	return seen[rowIndex]
}

func maxProfit(prices []int) int {
	var maxProfit int
	mem := map[int]int{}

	for i := 0; i < len(prices)-1; i++ {
		var localMax int
		if _, ok := mem[prices[i]]; !ok {
			localMax = prices[i+1]
			for j := i + 1; j < len(prices); j++ {
				if prices[j] > localMax {
					localMax = prices[j]
				}
			}
			mem[prices[i]] = localMax
		}
		if mem[prices[i]]-prices[i] > maxProfit {
			maxProfit = mem[prices[i]] - prices[i]
		}
	}
	return maxProfit
}
