package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
)

type Direction uint8

const (
	N Direction = iota
	E
	S
	W
)

type P struct{ x, y int }

var delta = map[Direction]P{
	N: {-1, 0},
	E: {0, 1},
	S: {1, 0},
	W: {0, -1},
}

func (p P) Step(dir Direction) P {
	return P{p.x + delta[dir].x, p.y + delta[dir].y}
}

type Grid struct {
	grid          []string
	Start, End    P
	Width, Height int
}

func (g Grid) OOB(p P) bool {
	return p.x < 0 || p.x >= len(g.grid) || p.y < 0 || p.y >= len(g.grid[0])
}

func (g Grid) Get(p P) uint8 {
	return g.grid[p.x][p.y]
}

func parse() Grid {
	var g Grid
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		g.grid = append(g.grid, scanner.Text())
	}
	g.Height = len(g.grid)
	g.Width = len(g.grid[0])
	g.Start = P{0, 1}
	g.End = P{g.Height - 1, g.Width - 2}
	return g
}

func main() {
	grid := parse()
	nodes := findNodes(grid)

	graph := NewGraph()
	graph.AddNode(grid.Start)
	graph.AddNode(grid.End)
	for _, n := range nodes {
		graph.AddNode(n)
	}

	for a := range graph {
		for b := range graph {
			if a != b {
				dist := distance(a, b, grid, graph, 0, map[P]struct{}{}, IgnoreSlopes)
				if dist != -1 {
					graph.AddEdge(a, b, dist)
				}
			}
		}
	}

	fmt.Println(slices.Max(dfs(grid.Start, 0, graph, grid.End, map[P]struct{}{grid.Start: {}})))
}

func findNodes(grid Grid) []P {
	var nodes []P

	for i := 0; i < grid.Height; i++ {
		for j := 0; j < grid.Width; j++ {
			p := P{i, j}

			if grid.Get(p) == '#' {
				continue
			}

			adj := countAdjacent(grid, p)

			if adj > 2 {
				nodes = append(nodes, p)
			}
		}
	}
	return nodes
}

func countAdjacent(grid Grid, p P) int {
	count := 0
	for _, dir := range []Direction{N, E, S, W} {
		adj := p.Step(dir)
		if !grid.OOB(adj) && grid.Get(adj) != '#' {
			count++
		}
	}
	return count
}

func distance(from, to P, grid Grid, graph Graph, weight int, visited map[P]struct{}, ignoreSlopes bool) int {
	visited[from] = struct{}{}

	if from == to {
		return weight
	}

	maxDistance := -1
	for _, dir := range []Direction{N, E, S, W} {

		adj := from.Step(dir)

		if _, ok := visited[adj]; ok {
			continue
		}

		if grid.OOB(adj) {
			continue
		}

		dest := grid.Get(adj)

		if dest == '#' {
			continue
		}

		if !ignoreSlopes {
			if dest == 'v' && dir != S {
				continue
			}

			if dest == '^' && dir != N {
				continue
			}

			if dest == '>' && dir != E {
				continue
			}

			if dest == '<' && dir != W {
				continue
			}
		}

		if !graph.IsNode(adj) {
			visited := maps.Clone(visited)
			distance := distance(adj, to, grid, graph, weight+1, visited, ignoreSlopes)
			if distance > maxDistance {
				maxDistance = distance
			}
			continue
		}

		if adj == to {
			distance := weight + 1
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}

	return maxDistance
}

func dfs(start P, score int, graph Graph, goal P, visited map[P]struct{}) []int {

	if start == goal {
		return []int{score}
	}

	var scores []int

	for _, edge := range graph.GetEdges(start) {
		if _, ok := visited[edge.To]; !ok {
			visited[edge.To] = struct{}{}
			scores = append(scores, dfs(edge.To, score+edge.Weight, graph, goal, visited)...)
			delete(visited, edge.To)
		}
	}

	return scores
}
