package main

type Node struct {
	Val int
}

type Graph struct {
	Nodes []*Node
	Edges map[Node]map[Node]struct{}
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	if _, ok := g.Edges[*n1]; !ok {
		g.Edges[*n1] = make(map[Node]struct{})
	}
	g.Edges[*n1][*n2] = struct{}{}

	if _, ok := g.Edges[*n2]; !ok {
		g.Edges[*n2] = make(map[Node]struct{})
	}
	g.Edges[*n2][*n1] = struct{}{}
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make([]*Node, 0),
		Edges: make(map[Node]map[Node]struct{}),
	}
}

// entry point for Depth First Search
func (g *Graph) DFS(start Node) <-chan Node {
	visited := make(map[Node]struct{})
	res := make(chan Node)

	go func() {
		g.dfs(visited, start, res)
		close(res)
	}()

	return res
}

// private recursive Depth First Search
func (g *Graph) dfs(visited map[Node]struct{}, n Node, res chan<- Node) {
	visited[n] = struct{}{}
	res <- n

	for k := range g.Edges[n] {
		if _, ok := visited[k]; !ok {
			g.dfs(visited, k, res)
		}
	}
}

// entry point for Breadth First Search
func (g *Graph) BFS(start Node) <-chan Node {
	ch := make(chan Node)
	go func() {
		toVisit := []Node{start}
		visited := map[Node]struct{}{start: {}}
		ch <- start

		for len(toVisit) > 0 {
			elem := toVisit[0]
			toVisit = toVisit[1:]

			for k := range g.Edges[elem] {
				if _, ok := visited[k]; !ok {
					visited[k] = struct{}{}
					ch <- k
					toVisit = append(toVisit, k)
				}
			}
		}
		close(ch)
	}()
	return ch
}
