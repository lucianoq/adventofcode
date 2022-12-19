package main

import "fmt"

const MinuteLimit = 32

func main() {
	blueprints := parse()

	mul := 1
	for _, b := range blueprints[:3] {
		mul *= bfs(b)
	}
	fmt.Println(mul)
}
