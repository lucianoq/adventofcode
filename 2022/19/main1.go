package main

import "fmt"

const MinuteLimit = 24

func main() {
	blueprints := parse()

	quality := 0
	for _, b := range blueprints {
		quality += int(b.ID) * bfs(b)
	}
	fmt.Println(quality)
}
