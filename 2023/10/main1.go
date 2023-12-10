package main

import "fmt"

func main() {
	m, start := parse()

	s1 := State{start.W(), E}
	s2 := State{start.E(), W}

	var steps int
	for steps = 1; s1.Current != s2.Current; steps++ {
		s1 = s1.Next(m)
		s2 = s2.Next(m)
	}

	fmt.Println(steps)
}
