package main

import (
	"bufio"
	"fmt"
	"os"
)

type P struct{ X, Y int }

func (p P) Neighbours() [4]P {
	return [4]P{
		{p.X, p.Y + 1}, // N
		{p.X + 1, p.Y}, // E
		{p.X, p.Y - 1}, // S
		{p.X - 1, p.Y}, // W
	}
}

type Status struct {
	Point  P
	Keybag int32
}

type Vault struct {
	Grid     map[P]string
	XMax     int
	YMax     int
	Entrance P
	NumKeys  int
}

func main() {
	vault := parse()
	fmt.Println(vault.FindKeysBFS())
}

func (v *Vault) FindKeysBFS() int {
	discovered := map[Status]struct{}{}
	toDo := []Status{} // let Q be a queue

	discovered[Status{v.Entrance, 0}] = struct{}{} // label root as discovered
	toDo = append(toDo, Status{v.Entrance, 0})     // Q.enqueue(root)

	depth := 0
	var curr Status

	for len(toDo) > 0 {

		for levelSize := len(toDo); levelSize > 0; levelSize-- {

			curr, toDo = toDo[0], toDo[1:]

			// if all keybag bits are turned on
			if curr.Keybag&(1<<v.NumKeys-1) == 1<<v.NumKeys-1 {
				return depth
			}

			for _, adj := range curr.Point.Neighbours() {

				cell := v.Grid[adj]

				// adjacent is wall
				if cell == "#" {
					continue // skip wall
				}

				// a door is found
				if cell[0] >= 'A' && cell[0] <= 'Z' {
					if curr.Keybag&(1<<(cell[0]-'A')) == 0 {
						continue // skip if we have no key for it
					}
				}

				keyBag := curr.Keybag

				// a key is found
				if cell[0] >= 'a' && cell[0] <= 'z' {
					// add it to the bitset
					keyBag |= 1 << (cell[0] - 'a')
				}

				nextStatus := Status{adj, keyBag}

				if _, ok := discovered[nextStatus]; !ok {
					discovered[nextStatus] = struct{}{}
					toDo = append(toDo, nextStatus)
				}
			}
		}

		depth++
	}

	return 0
}

func parse() *Vault {

	vault := &Vault{
		Grid:     map[P]string{},
		Entrance: P{},
	}

	numKeys := 0

	i := 0
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); i++ {
		for j, c := range scanner.Text() {
			switch {

			case c == '@': // Entrance
				vault.Entrance.X = i
				vault.Entrance.Y = j
				vault.Grid[P{i, j}] = "@"

			case c == '#': // wall
				vault.Grid[P{i, j}] = "#"

			case c == '.': // open
				vault.Grid[P{i, j}] = "."

			case c >= 'a' && c <= 'z': // key
				vault.Grid[P{i, j}] = string(c)
				numKeys++

			case c >= 'A' && c <= 'Z': // door
				vault.Grid[P{i, j}] = string(c)
			}

			vault.YMax = j + 1
		}
	}

	vault.XMax = i
	vault.NumKeys = numKeys

	return vault
}
