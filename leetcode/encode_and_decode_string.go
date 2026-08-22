package main

import (
	"strconv"
)

type Solution struct{}

// shayan hello something
// 0, len(shayan), len(shayan) + len(hello)

// this will give this n n1,n2,...nM, str1,str2,...,strM where nI = len(strI)
func (s *Solution) Encode(strs []string) string {
	var encoded string
	for _, str := range strs {
		encoded += strconv.Itoa(len(str)-1) + ";" + str
	}
	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	var decoded []string
	var tempCount, tempStr string
	var curr byte
	var currCount int
	var nFlag bool = true

	for i := 0; i < len(encoded); i++ {
		curr = encoded[i]
		if nFlag {
			if curr == ';' {
				nFlag = false
				currCount, _ = strconv.Atoi(tempCount)
				tempCount = ""
				continue
			}
			tempCount += string(curr)
		}
		if !nFlag {
			tempStr += string(curr)
			if currCount == 0 {
				decoded = append(decoded, tempStr)
				tempStr = ""
				nFlag = true
			}
			currCount--
		}
	}
	return decoded
}
