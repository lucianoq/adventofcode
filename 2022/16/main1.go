package main

import "fmt"

const MaxMinutes = 30

func main() {
	valves = parse()
	distances = createDistancesMatrix(valves)

	// Filter out valves with flow==0
	toOpen := []string{}
	for k, v := range valves {
		if v.Flow > 0 {
			toOpen = append(toOpen, k)
		}
	}

	fmt.Println(dfs(0, 0, 0, "AA", toOpen))
}
