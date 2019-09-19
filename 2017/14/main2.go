package main

import "fmt"

type Coord struct{ X, Y int }

func main() {
	g := NewGrid()

	fmt.Println(g.NumRegions())
}

func (g *Grid) NumRegions() int {
	var visited = make(map[Coord]struct{})

	region := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if _, ok := visited[Coord{i, j}]; !ok && g[i][j] {
				region++
				g.visitRegion(visited, Coord{i, j})
			}
		}
	}
	return region
}

func (g *Grid) visitRegion(visited map[Coord]struct{}, start Coord) {
	visited[start] = struct{}{}
	toDo := []Coord{start}

	for len(toDo) > 0 {
		elem := toDo[0]
		toDo = toDo[1:]

		for _, c := range g.adjacent(elem) {
			if _, ok := visited[c]; !ok {
				visited[c] = struct{}{}
				toDo = append(toDo, c)
			}
		}
	}
}

func (g *Grid) adjacent(c Coord) []Coord {
	res := make([]Coord, 0)
	for _, offset := range []Coord{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		x, y := c.X+offset.X, c.Y+offset.Y
		if x >= 0 && y >= 0 && x < size && y < size {
			if g[x][y] {
				res = append(res, Coord{x, y})
			}
		}
	}
	return res
}
