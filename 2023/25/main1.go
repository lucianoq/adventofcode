package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Graph map[string]map[string]struct{}

func (g Graph) Link(a, b string) {
	if _, ok := g[a]; !ok {
		g[a] = make(map[string]struct{})
	}
	g[a][b] = struct{}{}

	if _, ok := g[b]; !ok {
		g[b] = make(map[string]struct{})
	}
	g[b][a] = struct{}{}
}

func (g Graph) Unlink(a, b string) {
	delete(g[a], b)
	delete(g[b], a)
}

func (g Graph) Dot() string {
	sb := strings.Builder{}
	sb.WriteString("graph G {\n")
	for a := range g {
		for b := range g[a] {
			sb.WriteString(fmt.Sprintf("\t%s -- %s\n", a, b))
		}
	}
	sb.WriteString("}\n")
	return sb.String()
}

func parse() Graph {
	g := Graph{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		ff := strings.Split(line, ":")
		left := ff[0]
		right := strings.Fields(ff[1])

		for _, r := range right {
			g.Link(left, r)
		}
	}
	return g
}

func main() {
	g := parse()

	// Uncomment the following line to print the .dot file.
	//fmt.Println(g.Dot())
	// Use `neato -T png graph.dot >graph.png`
	// to generate the image included in the dir.

	g.Unlink("vps", "pzc")
	g.Unlink("xvk", "sgc")
	g.Unlink("dph", "cvx")

	fmt.Println(bfs(g, "vps") * bfs(g, "pzc"))
}

func bfs(g Graph, start string) int {
	visited := map[string]struct{}{}
	toDo := []string{start}
	var curr string
	for len(toDo) > 0 {
		curr, toDo = toDo[0], toDo[1:]
		for node := range g[curr] {
			if _, ok := visited[node]; !ok {
				visited[node] = struct{}{}
				toDo = append(toDo, node)
			}
		}
	}
	return len(visited)
}
