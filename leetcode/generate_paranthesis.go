package main

import "fmt"

func generateParenthesis(n int) []string {
	var res []string

	var bt func(curr string, open int, close int)
	bt = func(curr string, open int, close int) {
		fmt.Println(curr)
		if len(curr) == 2*n {
			res = append(res, curr)
			fmt.Println("done")
			return
		}

		if open < n { // just becuase we need to have open parenthesis?
			bt(curr+"(", open+1, close)
		}

		if close < open {
			bt(curr+")", open, close+1)
		}
	}

	bt("", 0, 0)
	return res
}
