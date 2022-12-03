package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		set := map[rune]struct{}{}
		dups := map[rune]struct{}{}

		for i, c := range line {
			if i < len(line)/2 {
				set[c] = struct{}{}
			} else {
				if _, ok := set[c]; ok {
					dups[c] = struct{}{}
				}
			}
		}

		for c := range dups {
			sum += priority(c)
		}
	}
	fmt.Println(sum)
}

func priority(r rune) int {
	if r <= 'Z' {
		return int(r) - 'A' + 26 + 1
	}
	return int(r) - 'a' + 1
}
