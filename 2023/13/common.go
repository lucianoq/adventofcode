package main

import (
	"io"
	"os"
	"strings"
)

type Pattern []string

func (p Pattern) Transpose() Pattern {
	newP := Pattern{}
	for c := 0; c < len(p[0]); c++ {
		var row string
		for r := 0; r < len(p); r++ {
			row += string(p[r][c])
		}
		newP = append(newP, row)
	}
	return newP
}

func levenshtein(s1, s2 string) int {
	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}
	return diff
}

func parse() []Pattern {
	body, _ := io.ReadAll(os.Stdin)
	patterns := []Pattern{}
	for _, c := range strings.Split(string(body), "\n\n") {
		patterns = append(patterns, strings.Split(strings.TrimSpace(c), "\n"))
	}
	return patterns
}
