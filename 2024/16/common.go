package main

import (
	"bufio"
	"container/heap"
	"math"
	"os"
)

type P struct{ x, y int }

type Direction uint8

func (d Direction) Clockwise() Direction        { return (d + 4 + 1) % 4 }
func (d Direction) Counterclockwise() Direction { return (d + 4 - 1) % 4 }

const (
	E Direction = iota
	S
	W
	N
)

var Delta = map[Direction]P{N: {-1, 0}, E: {0, 1}, S: {1, 0}, W: {0, -1}}

func parseInput() (map[P]struct{}, P, P) {
	var (
		m     = map[P]struct{}{}
		start P
		end   P
	)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		for j, ch := range scanner.Text() {
			switch ch {
			case 'S':
				m[P{i, j}] = struct{}{}
				start = P{i, j}
			case 'E':
				m[P{i, j}] = struct{}{}
				end = P{i, j}
			case '.':
				m[P{i, j}] = struct{}{}
			case '#':
			}
		}
	}
	return m, start, end
}

type Node struct {
	Pos       P
	Direction Direction
}

func Dijkstra(m map[P]struct{}, start, end P) (int, map[Node][]Node) {
	var (
		minScore = math.MaxInt64
		parent   = map[Node][]Node{}
	)

	startNode := Node{start, E}

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, pqNode{startNode, 0})

	scores := map[Node]int{startNode: 0}

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(pqNode)

		// we reached the end
		if curr.Pos == end {

			// stop only when we got over the minScore, as
			// we are interested into all the minScore paths
			if curr.Score > minScore {
				return minScore, parent
			}
			minScore = curr.Score
		}

		for _, n := range neighbours(m, curr) {
			if _, ok := scores[n.Node]; !ok || n.Score <= scores[n.Node] {
				parent[n.Node] = append(parent[n.Node], curr.Node)
				scores[n.Node] = n.Score
				heap.Push(pq, n)
			}
		}
	}
	panic("no path found")
}

func neighbours(m map[P]struct{}, curr pqNode) []pqNode {
	list := []pqNode{
		{Node{curr.Pos, curr.Direction.Clockwise()}, curr.Score + 1000},
		{Node{curr.Pos, curr.Direction.Counterclockwise()}, curr.Score + 1000},
	}

	delta := Delta[curr.Direction]
	if _, ok := m[P{curr.Pos.x + delta.x, curr.Pos.y + delta.y}]; ok {
		list = append(list, pqNode{Node{
			Pos:       P{curr.Pos.x + delta.x, curr.Pos.y + delta.y},
			Direction: curr.Direction,
		}, curr.Score + 1})
	}

	return list
}
