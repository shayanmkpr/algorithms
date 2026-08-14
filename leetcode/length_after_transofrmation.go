package main

import (
	"fmt"
)

/*
if its z, then replace it with ab, if its anything else, replace it with char++
since we want to do this with t transormations, we will keep the state of the app.
*/

func lengthAfterTransformations(s string, t int) int {
	runes := []rune(s)
	for range t {
		for i := 0; i < len(runes); i++ {
			fmt.Println(i, string(runes))
			if runes[i] != 'z' {
				runes[i] = runes[i] + 1
				continue
			} else {
				addedPart := make([]rune, len(runes[i+1:]))
				copy(addedPart, runes[i+1:])
				runes = append(append(runes[:i], []rune{'a', 'b'}...), addedPart...)
				i++
				continue
			}
		}
		fmt.Println(string(runes))
	}
	return len(runes)
}
