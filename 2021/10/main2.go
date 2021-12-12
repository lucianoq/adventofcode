package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var pointsIncomplete = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scores := []int{}

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

			// There is a wrong closing character
			// -> corrupted line, skip it
			continue Line
		}

		if len(stack) > 0 {
			// line incomplete
			score := 0
			for i := len(stack) - 1; i >= 0; i-- {
				score = 5*score + pointsIncomplete[stack[i]]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
