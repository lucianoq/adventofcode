package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type P struct{ x, y int }

var numKeypadMap = map[byte]P{
	'7': {0, 0}, '8': {0, 1}, '9': {0, 2},
	'4': {1, 0}, '5': {1, 1}, '6': {1, 2},
	'1': {2, 0}, '2': {2, 1}, '3': {2, 2},
	'_': {3, 0}, '0': {3, 1}, 'A': {3, 2},
}

var dirKeypadMap = map[byte]P{
	'_': {0, 0}, '^': {0, 1}, 'A': {0, 2},
	'<': {1, 0}, 'v': {1, 1}, '>': {1, 2},
}

func main() {
	codes := parseInput()
	sum := 0
	for _, code := range codes {
		directions := pressNumbers(code)
		seqLength := press(directions, Depth)

		numericPart, _ := strconv.Atoi(code[:len(code)-1])

		sum += seqLength * numericPart
	}
	fmt.Println(sum)
}

// type for the key in the map for press() memoization
type key = struct {
	seq   string
	depth int
}

// memoization for press()
var cachePress = map[key]int{}

// press returns the length of the directional string
// after applying `depth` times the indirection
func press(code string, depth int) int {
	if depth == 0 {
		return len(code)
	}

	// check in cache
	if size, ok := cachePress[key{code, depth}]; ok {
		return size
	}

	totalSize := 0
	for _, chunk := range strings.SplitAfter(code, "A") {
		totalSize += press(pressDirections(chunk), depth-1)
	}

	// save into cache
	cachePress[key{code, depth}] = totalSize

	return totalSize
}

func pressDirections(dir string) string {
	hole := dirKeypadMap['_']
	curr := dirKeypadMap['A']

	seq := ""
	for _, c := range dir {
		goal := dirKeypadMap[byte(c)]

		deltaH, deltaV := goal.y-curr.y, goal.x-curr.x
		hs, vs := getMovements(deltaH, deltaV)
		seq += prioritize(hs, vs, curr, goal, hole)

		curr.x += deltaV
		curr.y += deltaH

		seq += "A"
	}
	return seq
}

func pressNumbers(code string) string {
	hole := numKeypadMap['_']
	curr := numKeypadMap['A']

	seq := ""
	for _, c := range code {
		goal := numKeypadMap[byte(c)]

		deltaH, deltaV := goal.y-curr.y, goal.x-curr.x
		hs, vs := getMovements(deltaH, deltaV)
		seq += prioritize(hs, vs, curr, goal, hole)

		curr.x += deltaV
		curr.y += deltaH

		seq += "A"
	}
	return seq
}

func getMovements(deltaH, deltaV int) (string, string) {
	var hs, vs string
	for deltaH > 0 {
		hs += ">"
		deltaH--
	}
	for deltaH < 0 {
		hs += "<"
		deltaH++
	}
	for deltaV > 0 {
		vs += "v"
		deltaV--
	}
	for deltaV < 0 {
		vs += "^"
		deltaV++
	}
	return hs, vs
}

func prioritize(hs, vs string, from, to, avoid P) string {
	switch {
	case from.x == avoid.x && to.y == avoid.y:
		return vs + hs
	case from.y == avoid.y && to.x == avoid.x:
		return hs + vs
	case strings.Contains(hs, "<"):
		return hs + vs
	default:
		return vs + hs
	}
}

func parseInput() []string {
	var list []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	return list
}
