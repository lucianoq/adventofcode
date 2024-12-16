package main

import "fmt"

func main() {
	m, start, end := parseInput()
	_, parent := Dijkstra(m, start, end)

	cellVisited := map[P]struct{}{}
	dfs(parent, Node{end, N}, map[Node]struct{}{}, cellVisited)

	fmt.Println(len(cellVisited))
}

func dfs(parent map[Node][]Node, node Node, visited map[Node]struct{}, cellVisited map[P]struct{}) {
	visited[node] = struct{}{}
	cellVisited[node.Pos] = struct{}{}

	for _, n := range parent[node] {
		if _, ok := visited[n]; !ok {
			dfs(parent, n, visited, cellVisited)
		}
	}
}
