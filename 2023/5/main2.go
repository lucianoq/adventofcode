package main

import "fmt"

func main() {
	seeds, maps := parse()

	for loc := 0; ; loc++ {
		x := loc
		for m := len(maps) - 1; m >= 0; m-- {
			x = maps[m].InverseConvert(x)
		}

		for s := 0; s < len(seeds); s += 2 {
			if x >= seeds[s] && x <= seeds[s]+seeds[s+1] {
				fmt.Println(loc)
				return
			}
		}
	}
}

func (m Map) InverseConvert(to int) int {
	for _, t := range m {
		if to >= t.Destination && to <= t.Destination+t.Length {
			return t.Source + (to - t.Destination)
		}
	}
	return to
}
