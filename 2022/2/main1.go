package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	LosingPoints = iota * 3
	DrawingPoints
	WinningPoints
)

func main() {
	score := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ff := strings.Fields(scanner.Text())
		opponent := int(ff[0][0] - 'A')
		myPlay := int(ff[1][0] - 'X')
		score += myPlay + 1
		switch {
		case myPlay == (opponent+1)%3:
			score += WinningPoints
		case myPlay == opponent:
			score += DrawingPoints
		}
	}
	fmt.Println(score)
}
