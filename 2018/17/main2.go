package main

import "fmt"

func main() {
	g := parse()
	// g.Print()
	g.Fill()
	// g.Print()
	fmt.Println(g.CountStillWater())
}

func (g *Ground) CountStillWater() int {
	count := 0
	for j := g.MinY; j <= g.MaxY; j++ {
		for i := g.MinX - 100; i <= g.MaxX+100; i++ {
			p := g.Grid[C{i, j}]
			if p == SettledWater {
				count++
			}
		}
	}
	return count
}
