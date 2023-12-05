package main

import "fmt"

func main() {
	seeds, maps := parse()

	minLocation := 1<<31 - 1
	for _, s := range seeds {
		for _, m := range maps {
			s = m.Convert(s)
		}

		if s < minLocation {
			minLocation = s
		}
	}

	fmt.Println(minLocation)
}

func (m Map) Convert(from int) int {
	for _, t := range m {
		if from >= t.Source && from <= t.Source+t.Length {
			return t.Destination + (from - t.Source)
		}
	}
	return from
}
