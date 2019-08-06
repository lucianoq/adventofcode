package main

import (
	"bufio"
	"os"
	"strings"
)

func parse() []map[rune]int {
	freq := make([]map[rune]int, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		for i, c := range line {
			if i >= len(freq) {
				freq = append(freq, make(map[rune]int))
			}
			freq[i][c]++
		}
	}

	return freq
}
