package main

import "fmt"

func main() {
	g := parse()

	fmt.Println(count(g.DFS(Node{0})))
	//fmt.Println(count(g.BFS(Node{0})))
}

func count(ch <-chan Node) int {
	c := 0
	for range ch {
		c++
	}
	return c
}
