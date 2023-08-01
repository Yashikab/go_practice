package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	split_str := strings.Fields(s)
	wc := make(map[string]int)
	for _, s := range split_str {
		wc[s] += 1
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
