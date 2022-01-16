package main

import (
	"container/heap"
	_ "embed"
	"log"
)

const Cols = 13

var (
	//go:embed input
	Input string

	Energy           = map[Cell]int{Amber: 1, Bronze: 10, Copper: 100, Desert: 1000}
	ColumnByAmphipod = map[Cell]int{Amber: 3, Bronze: 5, Copper: 7, Desert: 9}
	AmphipodByColumn = map[int]Cell{3: Amber, 5: Bronze, 7: Copper, 9: Desert}
)

func Dijkstra(startGrid Grid) int {
	start := Node{startGrid, 0}

	frontier := &NodeHeap{}
	heap.Init(frontier)
	heap.Push(frontier, start)

	costSoFar := map[string]int{start.Grid.Hash(): 0}

	for frontier.Len() > 0 {

		curr := heap.Pop(frontier).(Node)

		if curr.Grid.Goal() {
			return curr.Energy
		}

		for _, n := range GenerateAllNextNodes(curr) {
			nextHash := n.Grid.Hash()
			if _, ok := costSoFar[nextHash]; !ok || n.Energy < costSoFar[nextHash] {
				costSoFar[nextHash] = n.Energy
				heap.Push(frontier, *n)
			}
		}
	}

	log.Fatal("no path found")
	return 0
}

func GenerateAllNextNodes(node Node) []*Node {
	nodes := []*Node{}
	for r := 0; r < Rows; r++ {
		for c := 0; c < Cols; c++ {
			switch node.Grid[r][c] {
			case Amber, Bronze, Copper, Desert:
				switch r {
				case 1:
					if n, ok := node.Grid.goHome(c); ok {
						n.Energy += node.Energy
						nodes = append(nodes, n)
					}
				default:
					if ns, ok := node.Grid.goOut(r, c); ok {
						for _, n := range ns {
							n.Energy += node.Energy
							nodes = append(nodes, n)
						}
					}
				}
			}
		}
	}
	return nodes
}
