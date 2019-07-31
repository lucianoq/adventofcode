package main

import "fmt"

const time = 2503

func main() {
	reindeer := parse()

	for t := 1; t < time; t++ {

		// find the highest distance that wins at this second
		max := 0
		for _, r := range reindeer {
			distance := r.Travelled(t)

			if distance > max {
				max = distance
			}
		}

		// find the winners and give them 1 point
		for _, r := range reindeer {
			if r.Travelled(t) == max {
				r.points++
			}
		}
	}

	// find the highest score
	max := 0
	for _, r := range reindeer {
		score := r.points
		if score > max {
			max = score
		}
	}

	fmt.Println(max)
}
