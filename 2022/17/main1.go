package main

import "fmt"

func main() {
	dir := getDir(parse())

	peak := 0
	for round := 0; round < 2022; round++ {
		peak = runRound(round, peak, dir)
	}
	fmt.Println(peak)
}
