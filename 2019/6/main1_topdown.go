package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	tree := parse()

	sum := 0
	for k := range tree {
		sum += findNumDescendants(tree, k)
	}

	fmt.Println(sum)
}

func parse() map[string][]string {
	tree := map[string][]string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nodes := strings.Split(scanner.Text(), ")")
		tree[nodes[0]] = append(tree[nodes[0]], nodes[1])
	}
	return tree
}

func findNumDescendants(tree map[string][]string, root string) int {
	var count int
	var current string
	toDo := []string{root}

	for len(toDo) > 0 {
		current, toDo = toDo[0], toDo[1:]

		for _, child := range tree[current] {
			count++
			toDo = append(toDo, child)
		}
	}

	return count
}
