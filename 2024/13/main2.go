package main

import "fmt"

func main() {
	configs := parseConfigs()

	for i := range configs {
		configs[i].P.x += 10000000000000
		configs[i].P.y += 10000000000000
	}

	tokens := 0
	for _, config := range configs {
		tokens += solve(config)
	}
	fmt.Println(tokens)
}
