package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		pos    = 50
		oldPos = 50
		count  = 0
	)

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		line := scanner.Text()
		steps, _ := strconv.Atoi(line[1:])

		// full loops of the command itself
		count += steps / 100

		// ignore the loops and keep only the meaningful steps
		steps = steps % 100

		if line[0] == 'L' {
			steps *= -1
		}

		oldPos = pos
		pos += steps

		switch {
		case pos == 0:
			count++

		case pos < 0:
			// If I'm starting from 0, I should not count the passage
			if oldPos != 0 {
				count++
			}
			pos += 100

		case pos > 99:
			count++
			pos -= 100
		}
	}

	fmt.Println(count)
}
