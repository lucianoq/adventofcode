package main

type Graph struct {
	Nodes map[char]*Node
}

type Node struct {
	ID    char
	Edges map[char]int
}

func (g *Graph) AddNode(n char) {
	if _, ok := g.Nodes[n]; ok {
		return
	}
	g.Nodes[n] = &Node{
		ID:    n,
		Edges: make(map[char]int),
	}
}

func (g *Graph) AddEdge(n1, n2 char, weight int) {
	if _, ok := g.Nodes[n1]; !ok {
		g.AddNode(n1)
	}

	oldWeight, ok := g.Nodes[n1].Edges[n2]
	if !ok || oldWeight > weight {
		g.Nodes[n1].Edges[n2] = weight
	}
}

func (g *Graph) GetNext(n char) map[char]int {
	return g.Nodes[n].Edges
}
