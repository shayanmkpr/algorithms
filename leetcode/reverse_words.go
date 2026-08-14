package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	vowel := map[rune]bool{}
	var vowelCount int
	fmt.Println(words)
	for _, ch := range []rune{'a', 'e', 'i', 'o', 'u'} {
		vowel[ch] = true
	}
	// find the number of vowels in the first word
	for _, ch := range words[0] {
		if vowel[ch] {
			vowelCount++
		}
	}

	for wordIdx, word := range words[1:] {
		count := 0
		for _, ch := range word {
			if vowel[ch] {
				count++
			}
		}
		if count == vowelCount {
			newWord := ""
			runes := []rune(word)
			for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}
			newWord = string(runes)
			words[wordIdx+1] = newWord
		}
	}
	var sb strings.Builder
	sb.WriteString(words[0])
	for _, w := range words[1:] {
		sb.WriteByte(' ')
		sb.WriteString(w)
	}
	return sb.String()
}
