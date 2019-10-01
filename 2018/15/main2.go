package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(FindOutcomeLowestAttack(os.Stdin))
}

func FindOutcomeLowestAttack(reader io.Reader) int {
	world, units := parse(reader)

	for attack := 38; ; attack++ {
		g := NewGame(world, units, attack)

		g.Run()

		if g.ElvesWonNoLosses() {
			return g.Outcome()
		}
	}
}

func (g *Game) End() bool {
	return g.NumElves < g.MaxElves || g.NumGoblins == 0
}
