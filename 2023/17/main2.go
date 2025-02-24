package main

import "fmt"

func main() {
	g := parse()

	fmt.Println(Dijkstra(g, 4, 10))
}
