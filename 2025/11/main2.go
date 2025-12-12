package main

import "fmt"

func main() {
	fmt.Println(dfs(parseInput(), "svr", "out", false, false))
}

type Args struct {
	start, goal string
	dac, fft    bool
}

var memoize = map[Args]int{}

func dfs(m map[string][]string, start, goal string, dac, fft bool) int {
	if v, ok := memoize[Args{start, goal, dac, fft}]; ok {
		return v
	}

	switch start {
	case "dac":
		dac = true
	case "fft":
		fft = true
	}

	if start == goal && dac && fft {
		return 1
	}

	sum := 0
	for _, o := range m[start] {
		sum += dfs(m, o, goal, dac, fft)
	}
	memoize[Args{start, goal, dac, fft}] = sum
	return sum
}
