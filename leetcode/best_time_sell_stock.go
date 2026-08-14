package main

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
