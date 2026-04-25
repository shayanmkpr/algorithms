package main

import "fmt"

func letterCombinations(digits string) []string {
	fmt.Println(digits)
	phoneMap := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
		'0': {" "},
	}
	var res []string
	var path []rune
	var backtracking func(pos int)
	backtracking = func(pos int) {
		if pos == len(digits) {
			res = append(res, string(path))
			return // why return? will return make us leave the function?
		}

		letters := phoneMap[rune(digits[pos])]

		for _, letter := range letters {
			path = append(path, []rune(letter)[0])
			backtracking(pos + 1)
			// basically the else
			path = path[:len(path)-1]
		}
	}

	backtracking(0)

	return res
}
