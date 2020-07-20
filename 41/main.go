package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var m map[string]int = map[string]int {}

	arr := strings.Fields(s)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
