package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	valves    map[string]Valve
	distances map[string]map[string]int
)

type Valve struct {
	Flow int
	Dest []string
}

func parse() map[string]Valve {
	valves := map[string]Valve{}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		var (
			name string
			flow int
		)
		fmt.Sscanf(line, "Valve %s has flow rate=%d; tunnels lead to valve", &name, &flow)
		dest := strings.TrimPrefix(strings.Split(line, "valve")[1], "s")
		valves[name] = Valve{
			Flow: flow,
			Dest: strings.Split(strings.TrimSpace(dest), ", "),
		}
	}
	return valves
}

func createDistancesMatrix(valves map[string]Valve) map[string]map[string]int {
	dist := map[string]map[string]int{}

	// init matrix to +Inf
	for v1 := range valves {
		dist[v1] = map[string]int{}
		for v2 := range valves {
			dist[v1][v2] = 1 << 30
		}
	}

	// set values from input
	for v1, vObj := range valves {
		for _, v2 := range vObj.Dest {
			dist[v1][v2] = 1
		}
	}

	// Floyd-Warshall Algorithm
	for k := range valves {
		for i := range valves {
			for j := range valves {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}

	return dist
}

func dfs(
	minute int,
	pressure int,
	flow int,
	pos string,
	toOpen []string,
) int {

	// project the result like we're done with opening valves
	max := pressure + (MaxMinutes-minute)*flow

	for _, v := range toOpen {

		// Need 1 to open the valve as well
		dist := distances[pos][v] + 1

		if minute+dist < MaxMinutes {
			if s := dfs(
				minute+dist,
				pressure+dist*flow,
				flow+valves[v].Flow,
				v,
				remove(toOpen, v),
			); s > max {
				max = s
			}
		}
	}
	return max
}

// less readable but more efficient than any reslice
// or single loop with append().
func remove(list []string, item string) []string {
	newList := make([]string, len(list)-1)
	i := 0
	for i = 0; i < len(list); i++ {
		if list[i] == item {
			break
		}
		newList[i] = list[i]
	}
	for i = i + 1; i < len(list); i++ {
		newList[i-1] = list[i]
	}
	return newList
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
