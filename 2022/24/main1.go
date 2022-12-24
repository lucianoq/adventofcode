package main

import "fmt"

func main() {
	parse()
	precomputeStates()
	fmt.Println(bfs(Entrance, Exit, 0))
}
