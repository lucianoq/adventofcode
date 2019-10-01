package main

import (
	"fmt"
)

type Game struct {
	World           map[C]bool
	Units           map[C]*Unit
	Rows, Cols      int
	NumElves        int
	NumGoblins      int
	MaxElves        int
	CompletedRounds int
	ElvesAttack     int
}

func NewGame(world map[C]bool, units map[C]Unit, elvesAttack int) *Game {
	maxRow, maxCol := 0, 0

	for c := range world {
		if c.X > maxRow {
			maxRow = c.X
		}
		if c.Y > maxCol {
			maxCol = c.Y
		}
	}

	g := &Game{
		World: world,
		Rows:  maxRow + 1,
		Cols:  maxCol + 1,
	}

	newUnits := make(map[C]*Unit, len(units))

	numElves, numGoblins := 0, 0

	for c, u := range units {
		newUnits[c] = &Unit{
			Type: u.Type,
			HP:   u.HP,
			AP:   u.AP,
			Pos:  u.Pos,
			Game: g,
		}
		if units[c].Type == Elf {
			newUnits[c].AP = elvesAttack
			numElves++
		} else {
			numGoblins++
		}
	}

	g.Units = newUnits
	g.NumElves = numElves
	g.NumGoblins = numGoblins
	g.MaxElves = numElves
	g.ElvesAttack = elvesAttack

	return g
}

func (g *Game) Run() int {
	for {
		if !g.Round() {
			break
		}
		g.CompletedRounds++
	}

	return g.Outcome()
}

// Returns round complete
func (g *Game) Round() bool {

	sortedUnits := g.SortedUnits()

	for _, u := range sortedUnits {

		if g.End() {
			return false
		}

		if u.HP <= 0 {
			continue
		}

		if !u.IsCloseEnemy() {
			u.Move()
		}

		if u.IsCloseEnemy() {
			u.Attack()
		}
	}

	return true
}

func (g *Game) UnitsAround(c C) []*Unit {
	res := make([]*Unit, 0)
	for _, c2 := range c.Near() {
		if u, ok := g.Units[c2]; ok {
			res = append(res, u)
		}
	}
	return res
}

func (g *Game) Empty(c C) bool {
	_, ok := g.Units[c]
	return !ok
}

func (g *Game) SortedUnits() []*Unit {
	res := make([]*Unit, 0)
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			if u, ok := g.Units[C{i, j}]; ok {
				res = append(res, u)
			}
		}
	}
	return res
}

func (g *Game) Outcome() int {
	sumHP := 0
	for _, u := range g.Units {
		sumHP += u.HP
	}
	return g.CompletedRounds * sumHP
}

func (g *Game) Print() {
	fmt.Println()

	fmt.Printf("Elves Attack: %d\n", g.ElvesAttack)
	fmt.Printf("Completed rounds: %d\n", g.CompletedRounds)

	fmt.Print("   ")
	for j := 0; j < g.Cols; j++ {
		fmt.Print(j % 10)
	}
	fmt.Println()
	for i := 0; i < g.Rows; i++ {

		fmt.Printf("%2d ", i)

		for j := 0; j < g.Cols; j++ {

			if u, ok := g.Units[C{i, j}]; ok {
				if u.Type == Elf {
					fmt.Print("E")
				} else {
					fmt.Print("G")
				}
				continue
			}

			if g.World[C{i, j}] {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("   ")

		for j := 0; j < g.Cols; j++ {
			if u, ok := g.Units[C{i, j}]; ok {
				if u.Type == Elf {
					fmt.Printf("E(%d), ", u.HP)
				} else {
					fmt.Printf("G(%d), ", u.HP)
				}
			}
		}

		fmt.Println()
	}
}

func (g *Game) ElvesWonNoLosses() bool {
	return g.NumElves == g.MaxElves && g.NumGoblins == 0
}

func (g *Game) bfs(c1, c2 C) (steps int, nextStep C) {
	// toDo := make([]C, 1, 10)
	// toDo[0] = c1
	toDo := []C{c1}
	discovered := map[C]struct{}{c1: {}}
	previous := map[C]C{c1: nullC}

	var found bool
	var current C

	for len(toDo) > 0 {
		current, toDo = toDo[0], toDo[1:]

		if current == c2 {
			found = true
			break
		}

		for _, cell := range current.Near() {

			// if it is a walkable cell
			if g.World[cell] {

				// if there is no-one
				if _, ok := g.Units[cell]; !ok {

					// if never discovered before
					if _, ok := discovered[cell]; !ok {
						discovered[cell] = struct{}{}
						toDo = append(toDo, cell)
						previous[cell] = current
					}
				}
			}
		}
	}

	// no path, no move
	if !found {
		return -1, nullC
	}

	for p := current; p != c1; p = previous[p] {
		nextStep = p
		steps++
	}
	return steps, nextStep
}
