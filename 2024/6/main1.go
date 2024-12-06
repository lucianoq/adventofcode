package main

import (
	"fmt"

	"github.com/lucianoq/container/set"
)

func main() {
	m, p, dir := parseMap()

	visited := set.Set[P]{}
	visited.Add(p)

	for {
		var next P

		switch dir {
		case N:
			next = P{p.x - 1, p.y}
		case E:
			next = P{p.x, p.y + 1}
		case S:
			next = P{p.x + 1, p.y}
		case W:
			next = P{p.x, p.y - 1}
		}

		if next.x < 0 || next.x >= Size || next.y < 0 || next.y >= Size {
			break
		}

		if _, ok := m[next]; ok {
			dir = (dir + 1) % 4
		} else {
			p = next
			visited.Add(p)
		}
	}

	fmt.Println(visited.Len())
}
