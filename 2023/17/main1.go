package main

import (
	"container/heap"
	"fmt"
)

// const Size = 141
const Size = 13

var (
	Start = P{0, 0}
	End   = P{12, 12}
)

func main() {
	m := parseInput()
	fmt.Println(Dijkstra(m, Start, End))
}

type P struct{ x, y int }

type Node struct {
	P
	Dir
}

func Dijkstra(m map[P]int, start, goal P) int {
	parent := map[P]P{}
	startNode := Node{start, E}
	heats := map[Node]int{startNode: 0}

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, pqNode{startNode, 0})

	for pq.Len() > 0 {

		curr := heap.Pop(pq).(pqNode)

		// we reached the end
		if curr.Node.P == goal {

			for x := goal; x != start; x = parent[x] {

			}

			return curr.Heat
		}

		for _, n := range neighbours(m, curr) {

			if _, ok := heats[n.Node]; !ok || n.Heat <= heats[n.Node] {
				parent[n.P] = parent[n.Node], curr.Node
				heats[n.Node] = n.Heat
				heap.Push(pq, n)
			}
		}
	}
	panic("no path found")
}

var Delta = map[Dir]P{
	N: {-1, 0},
	S: {1, 0},
	W: {0, -1},
	E: {0, 1},
}

func neighbours(m map[P]int, node pqNode) []pqNode {
	var neighbs []pqNode

	for _, d := range []Dir{node.Dir.Clock(), node.Dir.CounterClock()} {
		delta := Delta[d]
		heat := node.Heat
		for i := 0; i < 3; i++ {
			next := P{node.Node.x + i*delta.x, node.Node.y + i*delta.y}

			if next.x < 0 || next.x >= Size {
				continue
			}
			if next.y < 0 || next.y >= Size {
				continue
			}

			heat += m[next]
			neighbs = append(neighbs,
				pqNode{
					Node: Node{
						P:   next,
						Dir: d,
					},
					Heat: node.Heat + m[next],
				})
		}
	}
	return neighbs
}
