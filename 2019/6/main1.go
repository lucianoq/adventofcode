package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	parent := parse()

	sum := 0
	for k := range parent {
		sum += findNumAncestors(parent, k)
	}

	fmt.Println(sum)
}

func parse() map[string]string {
	parent := make(map[string]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nodes := strings.Split(scanner.Text(), ")")
		parent[nodes[1]] = nodes[0]
	}
	return parent
}

func findNumAncestors(parent map[string]string, node string) int {
	var count int
	for n := parent[node]; n != ""; n = parent[n] {
		count++
	}
	return count
}
