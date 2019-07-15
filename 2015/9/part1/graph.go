package main

import "math"

type Graph struct {
	nodes map[string]bool
	edges map[string][]Edge
}

type Edge struct {
	From     string
	To       string
	Distance int
}

func NewGraph() *Graph {
	g := &Graph{
		nodes: make(map[string]bool),
		edges: make(map[string][]Edge, 0),
	}

	return g
}

func (g *Graph) AddEdge(from, to string, d int) {
	g.nodes[from] = true
	g.nodes[to] = true
	g.edges[from] = append(g.edges[from], Edge{
		From:     from,
		To:       to,
		Distance: d,
	})
	g.edges[to] = append(g.edges[to], Edge{
		From:     to,
		To:       from,
		Distance: d,
	})
}

func (g *Graph) MinPath() int {
	paths := make([]int, 0)
	for n := range g.nodes {
		visited := make(map[string]bool)
		paths = append(paths, g.shortest(n, visited))
	}

	return min(paths)
}

func (g *Graph) shortest(from string, visited map[string]bool) int {
	visited[from] = true

	if len(visited) == len(g.nodes) {
		return 0
	}

	paths := make([]int, 0)
	for _, e := range g.edges[from] {
		if !visited[e.To] {
			paths = append(paths, e.Distance+g.shortest(e.To, deepCopy(visited)))
		}
	}

	return min(paths)
}

func deepCopy(orig map[string]bool) map[string]bool {
	newMap := make(map[string]bool)
	for k, v := range orig {
		newMap[k] = v
	}
	return newMap
}

func min(l []int) int {
	m := math.MaxInt64
	for _, x := range l {
		if x < m {
			m = x
		}
	}
	return m
}
