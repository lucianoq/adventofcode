package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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

func parse() []Reindeer {
	reindeer := make([]Reindeer, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		var r Reindeer

		n, _ := fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.",
			&r.name,
			&r.speed,
			&r.runTime,
			&r.restTime,
		)
		if n != 4 {
			log.Fatal("parsing failed")
		}

		reindeer = append(reindeer, r)
	}
	return reindeer
}

type Reindeer struct {
	name     string
	speed    int
	runTime  int
	restTime int
}

func (r Reindeer) Travelled(time int) int {
	period := r.runTime + r.restTime
	exact := time / period
	remaining := time % period
	return (exact*r.runTime + min(remaining, r.runTime)) * r.speed
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
