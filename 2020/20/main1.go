package main

import "fmt"

func main() {
	g := NewGraph()
	tiles := parse()

	for id1, tile1 := range tiles {
		for id2, tile2 := range tiles {
			if id1 != id2 {
				for s1 := Side(0); s1 < 8; s1++ {
					for s2 := Side(0); s2 < 8; s2++ {
						if Overlap(tile1.Edge(s1), tile2.Edge(s2)) {
							g.AddEdge(id1, id2)
						}
					}
				}
			}
		}
	}

	// corners are the nodes with just 2 edges

	mul := 1
	for n := range g.edges {
		if len(g.edges[n]) == 2 {
			mul *= n
		}
	}
	fmt.Println(mul)
}
