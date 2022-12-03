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

		set := map[byte]struct{}{}

		for i := 0; i < len(line)/2; i++ {
			set[line[i]] = struct{}{}
		}

		for i := len(line) / 2; i < len(line); i++ {
			if _, ok := set[line[i]]; ok {
				sum += priority(line[i])
				break
			}
		}
	}
	fmt.Println(sum)
}

func priority(c byte) int {
	if c <= 'Z' {
		return int(c) - 'A' + 26 + 1
	}
	return int(c) - 'a' + 1
}
