package main

import "fmt"

func main() {
	parse()
	precomputeStates()
	fmt.Println(bfs(Entrance, Exit, bfs(Exit, Entrance, bfs(Entrance, Exit, 0))))
}
