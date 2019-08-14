package main

import (
	"fmt"
	"math"
	"sort"
)

const salt = "yjdafjpo"

type Hash func(salt string, i int) string

func run(hashFn Hash) {
	var approved []int
	candidates := make(map[string][]int)
	limit := math.MaxInt64

	for i := 0; i <= limit; i++ {
		hex := hashFn(salt, i)

		for c := range quintuples(hex) {
			for _, candidateID := range candidates[c] {
				if i <= candidateID+1000 {
					approved = append(approved, candidateID)
					if len(approved) == 64 {
						// When I found at least 64 items, I need to be sure that
						// there aren't other items with a smaller ID and a validation
						// in the next 1000 IDs, so I run other 1000 iterations
						limit = i + 1000
					}
				}
			}
			candidates[c] = []int{}
		}

		if c, ok := triples(hex); ok {
			candidates[c] = append(candidates[c], i)
		}
	}

	// approved now could be greater than 64 items. I need to sort items and
	// take the 64th
	sort.Ints(approved)
	fmt.Println(approved[63])
}

func triples(s string) (string, bool) {
	for i := 0; i < len(s)-2; i++ {
		if same(s[i : i+3]) {
			return string(s[i]), true
		}
	}
	return "", false
}

func quintuples(s string) map[string]bool {
	res := make(map[string]bool)
	for i := 0; i < len(s)-4; i++ {
		if same(s[i : i+5]) {
			res[string(s[i])] = true
			i += 4
		}
	}
	return res
}

func same(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}
