package main

import "fmt"

const Rows = 5

func main() {
	grid := FromString(Input)
	fmt.Println(Dijkstra(grid))
}
