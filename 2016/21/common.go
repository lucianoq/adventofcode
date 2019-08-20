package main

import (
	"bufio"
	"os"
)

func parse() []string {
	scanner := bufio.NewScanner(os.Stdin)

	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
