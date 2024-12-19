package main

import (
	"fmt"
	"strings"
)

func main() {
	patterns, designs := parseInput()

	cache := map[string]int{}

	sum := 0
	for _, d := range designs {
		sum += ways(d, patterns, cache)
	}
	fmt.Println(sum)
}

func ways(d string, patterns []string, cache map[string]int) int {
	if len(d) == 0 {
		return 1
	}

	if w, ok := cache[d]; ok {
		return w
	}

	w := 0
	for _, p := range patterns {
		if strings.HasPrefix(d, p) {
			w += ways(d[len(p):], patterns, cache)
		}
	}

	cache[d] = w
	return w
}
