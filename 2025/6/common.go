package main

import (
	"bufio"
	"os"
)

var neutral = map[byte]int{
	'+': 0,
	'*': 1,
}

var fn = map[byte]func(int, int) int{
	'+': func(a, b int) int { return a + b },
	'*': func(a, b int) int { return a * b },
}

func parseInput() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
