package main

import (
	"fmt"
	"log"
)

type C struct{ X, Y int }

const (
	Rocky = iota
	Wet
	Narrow
)

const (
	Depth     = 8787
	TargetRow = 725
	TargetCol = 10
	// Depth     = 510
	// TargetRow = 10
	// TargetCol = 10
	Mod = 20183
)

func makeGridEL(overX, overY int) map[C]int {

	// We can prove that
	//   (a+b) % n = (a%n)+(b%n)
	//   (a⋅b) % n = (a%n)⋅(b%n)
	//
	// it follows that
	//   (geo * geo) + Depth % Mod
	// is equal to
	//   [(geo + Depth) % Mod] * [ ( geo + Depth ) % Mod ]
	//
	// so we can store Erosion Level directly instead of creating a grid
	// of geological indexes that are impossible to fit in memory.
	// Erosion level is always < Mod.

	mouth := C{0, 0}
	target := C{TargetCol, TargetRow}

	g := make(map[C]int)

	for y := 0; y <= TargetRow+overY; y++ {
		for x := 0; x <= TargetCol+overX; x++ {
			p := C{x, y}

			if p == mouth || p == target {
				g[p] = (0 + Depth) % Mod
				continue
			}

			if p.Y == 0 {
				g[p] = ((p.X * 16807) + Depth) % Mod
				continue
			}
			if p.X == 0 {
				g[p] = ((p.Y * 48271) + Depth) % Mod
				continue
			}

			// flowing from left to right and then from top to bottom
			// ensure that above and left cell have been already populated
			left := g[C{p.X - 1, p.Y}]
			top := g[C{p.X, p.Y - 1}]

			g[p] = ((left * top) + Depth) % Mod
		}
	}
	return g
}

func Print(g map[C]int, overX, overY int) {
	mouth := C{0, 0}
	target := C{TargetCol, TargetRow}

	for y := 0; y <= TargetRow+overY; y++ {
		for x := 0; x <= TargetCol+overX; x++ {
			p := C{x, y}

			if p == mouth {
				fmt.Print("M")
				continue
			}

			if p == target {
				fmt.Print("T")
				continue
			}

			t := g[p] % 3
			switch t {
			case Rocky:
				fmt.Print(".")
			case Narrow:
				fmt.Print("|")
			case Wet:
				fmt.Print("=")
			default:
				log.Fatal(t)
			}
		}
		fmt.Println()
	}
}
