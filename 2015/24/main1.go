package main

import "fmt"

func main() {
	packages, weight := parse()

	weightGoal := weight / 3

	ch := make(chan []int)
	go func() {
		// Wrote 6 after found out no groups smaller than 6 is possible.
		// TODO parametrise it
		generate(6, weightGoal, packages, []int{}, 0, ch)
		close(ch)
	}()

	fmt.Println(minEntanglement(packages, weightGoal, ch))
}
