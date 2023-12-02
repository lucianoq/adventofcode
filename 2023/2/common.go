package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Extraction struct {
	Red, Blue, Green int
}

func parse() map[int][]Extraction {
	scanner := bufio.NewScanner(os.Stdin)
	var db = map[int][]Extraction{}
	for scanner.Scan() {
		line := scanner.Text()
		game, ex := parseLine(line)
		db[game] = ex
	}
	return db
}

func parseLine(line string) (int, []Extraction) {
	ff := strings.Split(line, ":")
	var game int
	_, _ = fmt.Sscanf(ff[0], "Game %d", &game)

	extractions := []Extraction{}

	for _, ex := range strings.Split(ff[1], ";") {
		var extraction Extraction

		for _, c := range strings.Split(ex, ",") {
			var stones int
			var color string

			_, _ = fmt.Sscanf(c, "%d %s", &stones, &color)

			switch color {
			case "red":
				extraction.Red = stones
			case "blue":
				extraction.Blue = stones
			case "green":
				extraction.Green = stones
			}
		}
		extractions = append(extractions, extraction)
	}

	return game, extractions
}
