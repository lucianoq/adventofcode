package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var (
		pos   = 50
		count = 0
	)

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		line := scanner.Text()

		steps, _ := strconv.Atoi(line[1:])

		// ignore the loops and keep only the meaningful steps
		steps = steps % 100

		if line[0] == 'L' {
			steps *= -1
		}

		pos = (pos + steps + 100) % 100

		if pos == 0 {
			count++
		}
	}

	fmt.Println(count)
}
