package main

import (
	"fmt"

	"github.com/lucianoq/container/set"
)

func main() {
	m, p, dir := parseMap()

	var count int
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if testObstacle(m, P{i, j}, p, dir) {
				count++
			}
		}
	}

	fmt.Println(count)
}

type State struct {
	Pos P
	Dir uint8
}

func testObstacle(m map[P]any, obstacle, pos P, dir uint8) bool {

	// impossible if the guard is there
	if obstacle == pos {
		return false
	}

	// impossible if there is already an obstacle there
	if _, ok := m[obstacle]; ok {
		return false
	}

	m[obstacle] = struct{}{}
	defer func() {
		delete(m, obstacle)
	}()

	visited := set.Set[State]{}
	visited.Add(State{pos, dir})

	for {
		var next P

		switch dir {
		case N:
			next = P{pos.x - 1, pos.y}
		case E:
			next = P{pos.x, pos.y + 1}
		case S:
			next = P{pos.x + 1, pos.y}
		case W:
			next = P{pos.x, pos.y - 1}
		}

		if next.x < 0 || next.x >= Size || next.y < 0 || next.y >= Size {
			return false
		}

		if _, ok := m[next]; ok {
			dir = (dir + 1) % 4
		} else {
			pos = next
		}

		newState := State{pos, dir}
		if visited.Contains(newState) {
			return true
		}
		visited.Add(newState)
	}
}
