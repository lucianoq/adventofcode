package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dups := map[rune]int{}
	badges := []rune{}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		set := map[rune]struct{}{}
		for _, c := range line {
			set[c] = struct{}{}
		}

		for c := range set {
			dups[c]++
		}

		if i%3 == 2 {
			for c, freq := range dups {
				if freq == 3 {
					badges = append(badges, c)
				}
			}
			dups = map[rune]int{}
		}
	}

	sum := 0
	for _, b := range badges {
		sum += priority(b)
	}
	fmt.Println(sum)
}

func priority(r rune) int {
	if r <= 'Z' {
		return int(r) - 'A' + 26 + 1
	}
	return int(r) - 'a' + 1
}
