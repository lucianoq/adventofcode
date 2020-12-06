package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	count := 0
	m := map[rune]int{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			count += len(m)

			// reset
			m = map[rune]int{}
			continue
		}

		for _, c := range line {
			m[c]++
		}
	}

	// if file finishes without an empty line, we need to
	// handle also last group
	count += len(m)

	fmt.Println(count)
}
