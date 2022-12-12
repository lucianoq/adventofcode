package main

import "fmt"

func main() {
	min := 1<<63 - 1
	for x := 0; x < Height; x++ {
		for y := 0; y < Width; y++ {
			p := P{x, y}
			if Map[p] == 'a' {
				steps := bfs(p)
				if steps < min {
					min = steps
				}
			}
		}
	}
	fmt.Println(min)
}
