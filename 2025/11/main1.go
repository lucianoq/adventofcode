package main

import "fmt"

func main() {
	fmt.Println(dfs(parseInput(), "you", "out"))
}

func dfs(m map[string][]string, start, goal string) int {
	if start == goal {
		return 1
	}
	sum := 0
	for _, o := range m[start] {
		sum += dfs(m, o, goal)
	}
	return sum
}
