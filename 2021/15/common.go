package main

import (
	"container/heap"
	"fmt"
	"log"
)

type C struct{ x, y int }

var (
	Size = 100
	Dx   = [4]int{1, 0, -1, 0}
	Dy   = [4]int{0, 1, 0, -1}
)

func main() {
	grid := parse()
	source := C{0, 0}
	target := C{Size - 1, Size - 1}
	val := dijkstra(grid, source, target)
	fmt.Println(val)
}

func dijkstra(grid map[C]int, source, target C) int {
	dist := map[C]int{}

	explore := &NodeHeap{}

	for v := range grid {
		dist[v] = 1<<63 - 1
	}
	dist[source] = 0

	heap.Push(explore, Node{
		Pos:  source,
		Risk: dist[source],
	})

	for explore.Len() > 0 {
		u := heap.Pop(explore).(Node)

		if u.Pos == target {
			return dist[u.Pos]
		}

		for i := 0; i < 4; i++ {

			v := C{u.Pos.x + Dx[i], u.Pos.y + Dy[i]}

			// ignore out of bound
			if v.x < 0 || v.y < 0 || v.x >= Size || v.y >= Size {
				continue
			}

			alt := dist[u.Pos] + grid[v]

			if alt < dist[v] {
				dist[v] = alt
				heap.Push(explore, Node{v, dist[v]})
			}
		}
	}

	log.Fatal("no path found")
	return 0
}
