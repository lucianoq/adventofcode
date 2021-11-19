package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type C struct{ X, Y int }

func (c C) Down() C  { return C{c.X, c.Y + 1} }
func (c C) Left() C  { return C{c.X - 1, c.Y} }
func (c C) Right() C { return C{c.X + 1, c.Y} }
func (c C) Up() C    { return C{c.X, c.Y - 1} }

type Cell int

const (
	Empty Cell = iota
	Clay
	Water
	SettledWater
)

type Ground struct {
	Grid       map[C]Cell
	MaxX, MinX int
	MaxY, MinY int
}

func parse() *Ground {
	grid := make(map[C]Cell)

	r := regexp.MustCompile("^(\\w)=(\\d+), (\\w)=(\\d+)\\.\\.(\\d+)$")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		res := r.FindAllStringSubmatch(line, -1)

		a, err := strconv.Atoi(res[0][2])
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.Atoi(res[0][4])
		if err != nil {
			log.Fatal(err)
		}
		c, err := strconv.Atoi(res[0][5])
		if err != nil {
			log.Fatal(err)
		}

		switch res[0][1] {
		case "x":
			for y := b; y <= c; y++ {
				grid[C{a, y}] = Clay
			}
		case "y":
			for x := b; x <= c; x++ {
				grid[C{x, a}] = Clay
			}
		}
	}

	minX, minY := 1<<63-1, 1<<63-1
	maxX, maxY := 0, 0

	for k := range grid {
		if k.X > maxX {
			maxX = k.X
		}
		if k.X < minX {
			minX = k.X
		}
		if k.Y > maxY {
			maxY = k.Y
		}
		if k.Y < minY {
			minY = k.Y
		}
	}

	return &Ground{
		Grid: grid,
		MinY: minY,
		MaxY: maxY,
		MaxX: maxX,
		MinX: minX,
	}
}

func (g *Ground) Print() {
	fmt.Println()

	// Print could start from 0 and go 1 line more
	for y := 0; y <= g.MaxY+1; y++ {

		// Print also 10 left and 10 right
		for x := g.MinX - 3; x <= g.MaxX+3; x++ {
			switch g.Grid[C{x, y}] {
			case Empty:
				fmt.Print(" ")
			case Clay:
				fmt.Print("#")
			case Water:
				fmt.Print("|")
			case SettledWater:
				fmt.Print("~")
			}
		}
		fmt.Println()
	}
}

func (g *Ground) Fill() {
	toDo := []C{{500, 0}}
	var p C

	for len(toDo) > 0 {

		p, toDo = toDo[0], toDo[1:]

		// // Uncomment for animation
		// g.Print()
		// fmt.Println()
		// time.Sleep(300 * time.Millisecond)

		// If the point is out of range
		if p.Y > g.MaxY {
			// ignore point
			continue
		}

		// if point has been drawn by another spill, fill up
		if g.Grid[p] == SettledWater {
			toDo = append(toDo, p.Up())
			continue
		}

		// if the point has empty space below, go down
		down := p.Down()
		downCell := g.Grid[down]
		if downCell == Empty || downCell == Water {
			g.Grid[p] = Water
			toDo = append(toDo, down)
			continue
		}

		// no empty space below, time to spread

		// fill left
		leftMost := p
		for {
			curr, floor := g.Grid[leftMost], g.Grid[leftMost.Down()]
			if (curr == Empty || curr == Water) && (floor == Clay || floor == SettledWater) {
				g.Grid[leftMost] = Water
				leftMost = leftMost.Left()
				continue
			}
			break
		}

		// fill right
		rightMost := p
		for {
			curr, floor := g.Grid[rightMost], g.Grid[rightMost.Down()]
			if (curr == Empty || curr == Water) && (floor == Clay || floor == SettledWater) {
				g.Grid[rightMost] = Water
				rightMost = rightMost.Right()
				continue
			}
			break
		}

		if g.Grid[leftMost] == Clay && g.Grid[rightMost] == Clay {
			for i := leftMost.X + 1; i < rightMost.X; i++ {
				g.Grid[C{i, p.Y}] = SettledWater
			}
			toDo = append(toDo, p.Up())
			continue
		}

		if g.Grid[leftMost.Down()] == Empty {
			toDo = append(toDo, leftMost)
		}

		if g.Grid[rightMost.Down()] == Empty {
			toDo = append(toDo, rightMost)
		}
	}
}
