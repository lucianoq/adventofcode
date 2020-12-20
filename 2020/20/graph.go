package main

type Graph struct {
	edges map[int]map[int]struct{}
}

func (g *Graph) AddEdge(id1, id2 int) {
	if g.edges[id1] == nil {
		g.edges[id1] = map[int]struct{}{}
	}

	if g.edges[id2] == nil {
		g.edges[id2] = map[int]struct{}{}
	}

	g.edges[id1][id2] = struct{}{}
	g.edges[id2][id1] = struct{}{}
}

func NewGraph() *Graph {
	return &Graph{
		edges: map[int]map[int]struct{}{},
	}
}
