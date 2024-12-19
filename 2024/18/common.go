package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const Size = 70

type P struct{ x, y int }

var (
	Start = P{0, 0}
	Goal  = P{Size, Size}
	Delta = []P{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
)

func parseInput() []P {
	var m []P
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		ff := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(ff[0])
		y, _ := strconv.Atoi(ff[1])
		m = append(m, P{x, y})
	}
	return m
}

func BFS(corrupted map[P]struct{}) []P {
	queue := []P{Start}
	visited := map[P]struct{}{Start: {}}
	parent := map[P]P{}

	for len(queue) > 0 {
		var curr P
		curr, queue = queue[0], queue[1:]

		if curr == Goal {
			var path []P
			for p := Goal; p != Start; p = parent[p] {
				path = append(path, p)
			}
			return path
		}

		for _, d := range Delta {
			n := P{curr.x + d.x, curr.y + d.y}

			if n.x < 0 || n.x > Size || n.y < 0 || n.y > Size {
				continue
			}

			if _, ok := corrupted[n]; ok {
				continue
			}

			if _, ok := visited[n]; !ok {
				visited[n] = struct{}{}
				parent[n] = curr
				queue = append(queue, n)
			}
		}
	}

	return nil
}
