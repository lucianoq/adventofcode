package main

import "fmt"

func main() {
	m, pos, dir := parseMap()

	visited := run(m, pos, dir)

	fmt.Println(len(visited))
}
