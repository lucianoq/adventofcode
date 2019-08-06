package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var graph map[string][]string
var order string

func main() {
	graph = make(map[string][]string)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			var x, y string
			n, err := fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &x, &y)
			if err != nil || n != 2 {
				log.Fatal(err)
			}

			if graph[x] == nil {
				graph[x] = make([]string, 0)
			}
			if graph[y] == nil {
				graph[y] = make([]string, 0)
			}

			graph[y] = append(graph[y], x)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for next := findNext(); next != ""; next = findNext() {
		solve(next)
	}

	fmt.Println(order)

}

func findNext() string {
	next := make([]string, 0)
	for k, c := range graph {
		if len(c) == 0 {
			next = append(next, k)
		}
	}
	if len(next) == 0 {
		return ""
	}
	sort.Strings(next)
	return next[0]
}

func solve(n string) {
	order += n

	delete(graph, n)
	for k, c := range graph {
		del := -1
		for i := 0; i < len(c); i++ {
			if c[i] == n {
				del = i
				break
			}
		}
		if del != -1 {
			graph[k] = append(graph[k][:del], graph[k][del+1:]...)
		}
	}
}
