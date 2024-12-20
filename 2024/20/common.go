package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type P struct{ x, y int }

var Delta = []P{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func main() {
	walls, start, end := parseInput()

	path := bfs(walls, start, end)

	sum := 0
	for i := 0; i < len(path)-1; i++ {
		for j := i + 1; j < len(path); j++ {
			dist := manhattan(path[i], path[j])
			if dist > 0 && dist <= CheatTime {
				saved := j - i - dist
				if saved >= MinSavedTime {
					sum++
				}
			}
		}
	}
	fmt.Println(sum)
}

func parseInput() (map[P]struct{}, P, P) {
	walls := map[P]struct{}{}
	var start, end P
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		for j, ch := range scanner.Text() {
			switch ch {
			case '#':
				walls[P{i, j}] = struct{}{}
			case 'S':
				start = P{i, j}
			case 'E':
				end = P{i, j}
			}
		}
	}
	return walls, start, end
}

func bfs(walls map[P]struct{}, start, end P) []P {
	queue := []P{start}
	visited := map[P]struct{}{start: {}}
	parent := map[P]P{}

	for len(queue) > 0 {
		var curr P
		curr, queue = queue[0], queue[1:]

		if curr == end {
			return getPath(parent, start, end)
		}

		for _, d := range Delta {
			next := P{curr.x + d.x, curr.y + d.y}

			if _, ok := walls[next]; ok {
				continue
			}

			if _, ok := visited[next]; !ok {
				visited[next] = struct{}{}
				parent[next] = curr
				queue = append(queue, next)
			}
		}
	}
	return nil
}

func getPath(parent map[P]P, start, end P) []P {
	l := make([]P, 0)
	for p := end; p != start; p = parent[p] {
		l = append(l, p)
	}
	l = append(l, start)
	slices.Reverse(l)
	return l
}

func manhattan(p1, p2 P) int {
	return abs(p2.x-p1.x) + abs(p2.y-p1.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
