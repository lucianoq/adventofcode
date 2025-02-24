package main

import (
	"bufio"
	"os"
)

type Direction uint8

const (
	N Direction = iota
	E
	S
	W
)

type P struct{ x, y int }

var delta = map[Direction]P{
	N: {-1, 0},
	E: {0, 1},
	S: {1, 0},
	W: {0, -1},
}

func (p P) Step(dir Direction) P {
	return P{p.x + delta[dir].x, p.y + delta[dir].y}
}

type Grid struct {
	Heat [][]int
	Size int
}

func (g Grid) Contains(p P) bool {
	return p.x >= 0 && p.x < g.Size && p.y >= 0 && p.y < g.Size
}

func (g Grid) Goal(p P) bool {
	return p.x == g.Size-1 && p.y == g.Size-1
}

func (g Grid) HeatLoss(p P) int {
	return g.Heat[p.x][p.y]
}

func parse() Grid {
	scanner := bufio.NewScanner(os.Stdin)

	g := Grid{}

	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for c := 0; c < len(line); c++ {
			row = append(row, int(line[c]-'0'))
		}
		g.Heat = append(g.Heat, row)
	}
	g.Size = len(g.Heat)
	return g
}

func PossibleDirections(curr Direction, streak, minSteps, maxSteps int) []Direction {
	if streak < minSteps {
		return []Direction{
			curr, // straight
		}
	}

	if streak >= maxSteps {
		return []Direction{
			(curr + 1) % 4, // turn right
			(curr + 3) % 4, // turn left
		}
	}

	return []Direction{
		(curr + 1) % 4, // turn right
		curr,           // straight
		(curr + 3) % 4, // turn left
	}
}

type Node struct {
	Pos    P
	Dir    Direction
	Streak int
}

func Dijkstra(g Grid, minSteps, maxSteps int) int {

	// starting nodes
	startNode := Node{P{0, 0}, E, 0}
	startNode2 := Node{P{0, 0}, S, 0}
	visited := map[Node]int{startNode: 0, startNode2: 0}

	pq := NewPriorityQueue()
	pq.Insert(startNode, 0)
	pq.Insert(startNode2, 0)

	for pq.Len() > 0 {
		currentNode, currentHeatLoss := pq.PopMin()

		if g.Goal(currentNode.Pos) {
			return currentHeatLoss
		}

		for _, nextDir := range PossibleDirections(currentNode.Dir, currentNode.Streak, minSteps, maxSteps) {

			nextStreak := 1
			if nextDir == currentNode.Dir {
				nextStreak = currentNode.Streak + 1
			}

			nextPos := currentNode.Pos.Step(nextDir)

			// ignore points out of bound
			if !g.Contains(nextPos) {
				continue
			}

			nextNode := Node{
				Pos:    nextPos,
				Dir:    nextDir,
				Streak: nextStreak,
			}
			nextHeatLoss := currentHeatLoss + g.HeatLoss(nextPos)

			if oldHeatLoss, ok := visited[nextNode]; ok && oldHeatLoss <= nextHeatLoss {
				continue
			}

			visited[nextNode] = nextHeatLoss
			pq.Insert(nextNode, nextHeatLoss)
		}
	}

	return 0
}
