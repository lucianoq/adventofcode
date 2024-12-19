package main

import (
	"fmt"
	"strings"
)

func main() {
	patterns, designs := parseInput()

	cache := map[string]bool{}

	sum := 0
	for _, d := range designs {
		if possible(d, patterns, cache) {
			sum++
		}
	}
	fmt.Println(sum)
}

func possible(d string, patterns []string, cache map[string]bool) bool {
	if len(d) == 0 {
		return true
	}

	if t, ok := cache[d]; ok {
		return t
	}

	for _, p := range patterns {
		if strings.HasPrefix(d, p) {
			found := possible(d[len(p):], patterns, cache)
			if found {
				cache[d] = true
				return true
			}
		}
	}

	cache[d] = false
	return false
}
