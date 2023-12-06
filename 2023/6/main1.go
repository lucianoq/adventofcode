package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	times, distances := parse()

	mul := 1
	for i, raceTime := range times {
		count := 0
		for t := 0; t < raceTime; t++ {
			if (raceTime-t)*t > distances[i] {
				count++
			}
		}
		mul *= count
	}

	fmt.Println(mul)
}

func parse() ([]int, []int) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var times []int
	for _, t := range strings.Fields(scanner.Text())[1:] {
		n, _ := strconv.Atoi(t)
		times = append(times, n)
	}

	scanner.Scan()
	var distances []int
	for _, d := range strings.Fields(scanner.Text())[1:] {
		n, _ := strconv.Atoi(d)
		distances = append(distances, n)
	}

	return times, distances
}
