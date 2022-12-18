package main

import "fmt"

func main() {
	points := parse()

	surface := 0
	space := map[P]struct{}{}
	for p := range points {
		for _, a := range Adj {
			adj := p.Add(a)
			if _, ok := space[adj]; ok {
				surface--
			} else {
				surface++
			}
		}
		space[p] = struct{}{}
	}
	fmt.Println(surface)
}
