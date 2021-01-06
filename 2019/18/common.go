package main

import (
	"bufio"
	"os"
)

type char uint8

type P struct{ X, Y int8 }

func (p P) Neighbours() []P {
	return []P{
		{p.X - 1, p.Y}, // W
		{p.X, p.Y - 1}, // S
		{p.X + 1, p.Y}, // E
		{p.X, p.Y + 1}, // N
	}
}

type Vault struct {
	Grid    map[P]char
	XMax    int8
	YMax    int8
	POI     map[char]P
	NumKeys int
	Bots    []P
}

func (v *Vault) DFS(curr P, discovered map[P]int, depth int, results map[char]int) {
	discovered[curr] = depth
	for _, adj := range curr.Neighbours() {

		// ignore discovered
		if when, ok := discovered[adj]; ok && when <= depth {
			continue
		}

		cell, ok := v.Grid[adj]

		switch {

		// ignore walls
		case cell == '#', !ok:
			continue

		// if a door is found
		case cell >= 'A' && cell <= 'Z',

			// if a key is found
			cell >= 'a' && cell <= 'z':

			steps, ok := results[cell]
			if ok && steps <= depth {
				continue
			}

			results[cell] = depth + 1

		case cell == '.',
			cell >= '0' && cell <= '9':

			// go recursive
			v.DFS(adj, discovered, depth+1, results)
		}
	}
}

func parse() *Vault {
	vault := &Vault{
		Grid: map[P]char{},
		POI:  map[char]P{},
	}

	numKeys := 0

	i := int8(0)
	bots := char(0)
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); i++ {
		for idx, r := range scanner.Text() {
			c := char(r)
			j := int8(idx)

			switch {

			case c == '@': // Bot
				botName := bots + '0'
				vault.Bots = append(vault.Bots, P{i, j})
				vault.POI[botName] = P{i, j}
				bots++
				vault.Grid[P{i, j}] = botName

			case c == '#': // wall
				vault.Grid[P{i, j}] = '#'

			case c == '.': // open
				vault.Grid[P{i, j}] = '.'

			case c >= 'a' && c <= 'z': // key
				vault.POI[c] = P{i, j}
				vault.Grid[P{i, j}] = c
				numKeys++

			case c >= 'A' && c <= 'Z': // door
				vault.POI[c] = P{i, j}
				vault.Grid[P{i, j}] = c
			}

			vault.YMax = j + 1
		}
	}

	vault.XMax = i
	vault.NumKeys = numKeys

	return vault
}

func toGraph(v *Vault) *Graph {
	g := &Graph{Nodes: map[char]*Node{}}
	for id, p := range v.POI {
		g.AddNode(id)
		results := map[char]int{}
		v.DFS(p, map[P]int{}, 0, results)
		for goal, steps := range results {
			g.AddEdge(id, goal, steps)
		}
	}
	return g
}
