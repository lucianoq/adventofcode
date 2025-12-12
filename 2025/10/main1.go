package main

import "fmt"

func main() {
	list := parseInput()
	sum := 0
	for _, m := range list {
		sum += m.BFS()
	}
	fmt.Println(sum)
}

func (m Machine) BFS() int {
	if m.Indicator == 0 {
		return 0
	}

	dist := map[uint16]int{0: 0}
	queue := []uint16{0}
	for len(queue) > 0 {
		var curr uint16
		curr, queue = queue[0], queue[1:]
		currDist := dist[curr]

		for _, b := range m.Buttons {
			nextVal := curr ^ b

			if nextVal == m.Indicator {
				return currDist + 1
			}

			if _, ok := dist[nextVal]; !ok {
				dist[nextVal] = currDist + 1
				queue = append(queue, nextVal)
			}
		}
	}
	return -1
}
