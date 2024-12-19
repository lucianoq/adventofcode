package main

import (
	"bufio"
	"os"
	"strings"
)

func parseInput() ([]string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	patterns := strings.Split(scanner.Text(), ", ")
	scanner.Scan()
	designs := make([]string, 0)
	for scanner.Scan() {
		designs = append(designs, scanner.Text())
	}
	return patterns, designs
}
