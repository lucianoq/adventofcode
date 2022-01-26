package main

import "fmt"

func main() {
	aligned := alignAll(parse())

	set := map[Point]struct{}{}
	for _, s := range aligned {
		for _, b := range s.beacons {
			set[b] = struct{}{}
		}
	}

	fmt.Println(len(set))
}
