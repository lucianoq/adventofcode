package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
)

const MaxMinutes = 26

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	valves = parse()
	distances = createDistancesMatrix(valves)

	// Filter out valves with flow==0
	toOpen := []string{}
	for k, v := range valves {
		if v.Flow > 0 {
			toOpen = append(toOpen, k)
		}
	}

	maxPressure := 0
	for p := range partition(toOpen) {
		pressure := dfs(0, 0, 0, "AA", p[0]) + dfs(0, 0, 0, "AA", p[1])
		if pressure > maxPressure {
			maxPressure = pressure
		}
	}
	fmt.Println(maxPressure)
}

// Generate all possible partitions of a list of string
func partition(list []string) chan [2][]string {
	ch := make(chan [2][]string)
	go func() {
		// i+=2 will generate half of the partitions.
		// We skip those because they'd be already
		// calculated by their symmetrical.
		for i := uint64(0); i < 1<<len(list); i += 2 {
			part := [2][]string{}
			for j := 0; j < len(list); j++ {
				if i&(1<<j) != 0 {
					part[0] = append(part[0], list[j])
				} else {
					part[1] = append(part[1], list[j])
				}
			}
			ch <- part
		}
		close(ch)
	}()
	return ch
}
