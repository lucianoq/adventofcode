package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	Start = "start"
	End   = "end"
)

var paths int

func parse() map[string][]string {
	graph := map[string][]string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		ff := strings.Split(line, "-")

		graph[ff[0]] = append(graph[ff[0]], ff[1])
		graph[ff[1]] = append(graph[ff[1]], ff[0])
	}
	return graph
}
