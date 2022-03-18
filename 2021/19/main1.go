package main

import (
	"fmt"

	"github.com/lucianoq/container/set"
)

func main() {
	aligned := alignAll(parse())

	set := set.New[Point]()
	for _, s := range aligned {
		set.Add(s.beacons...)
	}

	fmt.Println(set.Len())
}
