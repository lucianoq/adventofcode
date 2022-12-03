package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sum := 0
	common := map[rune]int{}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		set := map[rune]struct{}{}
		for _, ch := range line {
			set[ch] = struct{}{}
		}

		for ch := range set {
			common[ch]++
		}

		if i%3 == 2 {
			for ch, freq := range common {
				if freq == 3 {
					sum += priority(ch)
					common = map[rune]int{}
					break
				}
			}
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
