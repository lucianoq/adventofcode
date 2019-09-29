package main

import "fmt"

func main() {
	v := Virus{
		Pos: C{0, 0},
		Dir: UP,
		G:   parse(),
	}

	for i := 0; i < 10000; i++ {
		v.Burst()
	}

	fmt.Println(v.BurstsCausedInfection)
}

func (s Status) Touch() Status {
	return (s + 2) % 4
}
