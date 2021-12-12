package main

import "fmt"

func main() {
	graph := parse()

	visited := map[string]int{}
	dfs(graph, visited, false, Start)

	fmt.Println(paths)
}

func dfs(graph map[string][]string, visited map[string]int, twiceTimeUsed bool, curr string) {
	if curr == End {
		paths++
		return
	}

	if curr[0] >= 'a' {
		visited[curr]++
	}

	for _, adj := range graph[curr] {

		if visited[adj] == 0 {
			dfs(graph, visited, twiceTimeUsed, adj)
			continue
		}

		if visited[adj] == 1 && !twiceTimeUsed && adj != Start {
			dfs(graph, visited, true, adj)
		}
	}

	if curr[0] >= 'a' {
		visited[curr]--
	}
}
