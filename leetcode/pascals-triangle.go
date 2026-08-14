package main

import "fmt"

func getRow(rowIndex int) []int {
	// var row [][]int
	row := [][]int{[]int{1}}
	mem := map[int][]int{}

	for i := 1; i <= rowIndex; i++ {
		fmt.Println("i", i)
		row = append(row, []int{})
		if _, ok := mem[i]; !ok {
			for j := 0; j <= i; j++ {
				row[i] = append(row[i], 0)
				fmt.Println("j", j)
				fmt.Println("row", row)
				if j == 0 || j == i {
					row[i][j] = 1
					continue
				}
				row[i][j] = row[i-1][j-1] + row[i-1][j]
			}
			mem[i] = row[i]
		}
	}
	return row[rowIndex]
}
