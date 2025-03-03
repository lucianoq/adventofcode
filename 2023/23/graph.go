package main

type Edge struct {
	To     P
	Weight int
}

type Graph map[P]map[Edge]struct{}

func NewGraph() Graph {
	return Graph{}
}

func (g Graph) AddNode(p P) {
	if _, ok := g[p]; !ok {
		g[p] = map[Edge]struct{}{}
	}
}

func (g Graph) IsNode(p P) bool {
	_, ok := g[p]
	return ok
}

func (g Graph) AddEdge(from, to P, weight int) {
	if _, ok := g[from]; !ok {
		g[from] = map[Edge]struct{}{}
	}
	if _, ok := g[to]; !ok {
		g[to] = map[Edge]struct{}{}
	}

	g[from][Edge{to, weight}] = struct{}{}
}

func (g Graph) GetEdges(from P) []Edge {
	list := []Edge{}
	for e := range g[from] {
		list = append(list, e)
	}
	return list
}
