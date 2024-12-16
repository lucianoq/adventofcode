package main

import "fmt"

func main() {
	score, _ := Dijkstra(parseInput())
	fmt.Println(score)
}
