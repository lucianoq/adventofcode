package main

import (
	"bufio"
	"fmt"
	"os"
)

var pointsCorrupted = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	score := 0

Line:
	for scanner.Scan() {
		line := scanner.Text()

		stack := []rune{}
		for _, c := range line {

			switch c {
			case '(', '[', '{', '<':
				stack = append(stack, c)
				continue
			}

			// It's a closing character
			lastOpen := stack[len(stack)-1]
			if c == lastOpen+1 || c == lastOpen+2 {
				stack = stack[:len(stack)-1]
				continue
			}

			// It's a wrong closing character
			// -> corrupted line
			score += pointsCorrupted[c]

			// ignore the remaining characters of the line
			continue Line
		}
	}

	fmt.Println(score)
}
