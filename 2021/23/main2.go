package main

import (
	"fmt"
	"strings"
)

const Rows = 7

func main() {
	grid := FromString(unfold(Input))
	fmt.Println(Dijkstra(grid))
}

func unfold(input string) string {
	lines := strings.Split(input, "\n")
	newLines := make([]string, 0, Rows)
	newLines = append(newLines, lines[:3]...)
	newLines = append(newLines, "  #D#C#B#A#")
	newLines = append(newLines, "  #D#B#A#C#")
	newLines = append(newLines, lines[3:]...)
	return strings.Join(newLines, "\n")
}
