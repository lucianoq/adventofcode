package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	count := 0
	words := 0
	m := make(map[rune]int)

	for scanner.Scan() {
		line := scanner.Text()

		// end of
		if line == "" {
			count += countEveryone(m, words)

			// reset
			m, words = make(map[rune]int), 0
			continue
		}

		words++
		for _, c := range line {
			m[c]++
		}
	}

	// if file finishes without an empty line, we need to
	// handle also last group
	if words > 0 {
		count += countEveryone(m, words)
	}

	fmt.Println(count)
}

func countEveryone(m map[rune]int, words int) int {
	count := 0
	for _, i := range m {
		if i == words {
			count++
		}
	}
	return count
}
