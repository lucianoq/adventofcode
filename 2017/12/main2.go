package main

import "fmt"

func main() {
	g := parse()

	groups := 0
	visited := map[Node]struct{}{}

	for _, n := range g.Nodes {

		if _, ok := visited[*n]; !ok {
			groups++

			ch := g.DFS(*n)
			//ch := g.BFS(*n)
			for n := range ch {
				visited[n] = struct{}{}
			}
		}
	}

	fmt.Println(groups)
}
