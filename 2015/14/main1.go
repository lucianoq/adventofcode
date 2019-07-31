package main

import "fmt"

const time = 2503

func main() {
	reindeer := parse()

	max := 0
	for _, r := range reindeer {
		distance := r.Travelled(time)

		if distance > max {
			max = distance
		}
	}

	fmt.Println(max)
}
