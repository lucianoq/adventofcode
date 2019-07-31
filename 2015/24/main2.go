package main

import "fmt"

func main() {

	packages, weight := parse()

	weightGoal := weight / 4

	ch := make(chan []int)
	go func() {
		// Wrote 4 after found out no groups smaller than 4 is possible.
		// TODO parametrise it
		generate(4, weightGoal, packages, []int{}, 0, ch)
		close(ch)
	}()

	fmt.Println(minEntanglement(packages, weightGoal, ch))
}
