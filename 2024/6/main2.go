package main

import "fmt"

func main() {
	m, pos, dir := parseMap()

	visited := run(m, pos, dir)

	var count int
	for obstacle := range visited {
		if testObstacle(m, obstacle, pos, dir) {
			count++
		}
	}
	fmt.Println(count)
}

type status struct {
	Pos P
	Dir Direction
}

func testObstacle(m map[P]struct{}, obstacle P, pos P, dir Direction) bool {

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

	visited := map[status]struct{}{
		status{pos, dir}: {},
	}

	for {
		next := pos.Next(dir)

		if next.x < 0 || next.x >= Size || next.y < 0 || next.y >= Size {
			return false
		}

		if _, ok := m[next]; ok {
			dir = (dir + 1) % 4
		} else {
			pos = next
		}

		st := status{pos, dir}
		if _, ok := visited[st]; ok {
			return true
		}
		visited[st] = struct{}{}
	}
}
