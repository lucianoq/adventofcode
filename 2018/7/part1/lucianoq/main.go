package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

 	"github.com/davecgh/go-spew/spew"
)

func main() {
	graph := make(map[string][]string)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		var x, y string
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &x, &y)

		if graph[x] == nil {
			graph[x] = make([]string, 0)
		}
		if graph[y] == nil {
			graph[y] = make([]string, 0)
		}
		graph[y] = append(graph[y], x)
	}

	spew.Dump(graph)
}
