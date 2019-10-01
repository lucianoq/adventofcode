package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(FindOutcome(os.Stdin))
}

func FindOutcome(reader io.Reader) int {
	world, units := parse(reader)

	g := NewGame(world, units, 3)

	g.Run()

	return g.Outcome()
}

func (g *Game) End() bool {
	return g.NumElves == 0 || g.NumGoblins == 0
}
