package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	strs := strings.Fields(s)
	for _, word := range strs{
		m[word]++
	}
	
	return m
}

func main() {
	wc.Test(WordCount)
}
