package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	all_strs := strings.Fields(s)
	the_map := make(map[string]int)

	for _, v := range all_strs {
		_, ok := the_map[v]
		if ok {
			the_map[v] = the_map[v] + 1
		} else {
			the_map[v] = 1
		}
	}

	return the_map
}

func main() {
	wc.Test(WordCount)
}
