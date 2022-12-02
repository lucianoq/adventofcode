package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	score := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ff := strings.Fields(scanner.Text())
		opponent := int(ff[0][0] - 'A')
		switch ff[1] {
		case "X":
			score += (opponent+2)%3 + 1
		case "Y":
			score += opponent + 1 + DrawingPoints
		case "Z":
			score += (opponent+1)%3 + 1 + WinningPoints
		}
	}
	fmt.Println(score)
}
