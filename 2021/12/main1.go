package main

import "fmt"

func main() {
	dfs(parse(), map[string]int{}, Start)
	fmt.Println(paths)
}

func dfs(graph map[string][]string, visited map[string]int, curr string) {
	if curr == End {
		paths++
		return
	}

	if curr[0] >= 'a' {
		visited[curr]++
	}

	for _, adj := range graph[curr] {
		if visited[adj] == 0 {
			dfs(graph, visited, adj)
		}
	}

	delete(visited, curr)
}
