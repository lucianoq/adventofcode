package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	maxCal := 0
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if sum > maxCal {
				maxCal = sum
			}
			sum = 0
			continue
		}
		cal, _ := strconv.Atoi(line)
		sum += cal
	}

	fmt.Println(maxCal)
}
