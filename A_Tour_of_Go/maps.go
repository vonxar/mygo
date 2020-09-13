package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordsCount := make(map[string]int)
	for i := 0; i < len(words); i++ {
		wordsCount[words[i]]++
	}
	return wordsCount
}

func main() {
	wc.Test(WordCount)
}
