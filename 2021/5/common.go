package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Point struct{ X, Y int }

type Segment struct{ P1, P2 Point }

func parse() []Segment {
	scanner := bufio.NewScanner(os.Stdin)
	segments := []Segment{}
	for scanner.Scan() {
		line := scanner.Text()

		ff := strings.Split(line, " -> ")
		start := strings.Split(ff[0], ",")
		end := strings.Split(ff[1], ",")
		x1, _ := strconv.Atoi(start[0])
		y1, _ := strconv.Atoi(start[1])
		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])

		segments = append(segments, Segment{Point{x1, y1}, Point{x2, y2}})
	}
	return segments
}

func minMax(x1, x2 int) (int, int) {
	if x1 <= x2 {
		return x1, x2
	}
	return x2, x1
}
