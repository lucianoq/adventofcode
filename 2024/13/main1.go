package main

import "fmt"

func main() {
	configs := parseConfigs()

	tokens := 0
	for _, config := range configs {
		tokens += solve(config)
	}
	fmt.Println(tokens)
}
