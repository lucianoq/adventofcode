package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parse() map[int]int {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()

	ff := strings.Split(strings.TrimSpace(line), ",")

	timers := map[int]int{}

	for _, f := range ff {
		num, _ := strconv.Atoi(f)

		timers[num]++
	}
	return timers
}

func evolve(timers map[int]int) map[int]int {
	newTimers := map[int]int{}

	newTimers[8] = timers[0]
	newTimers[6] = timers[0]
	for i := 1; i <= 8; i++ {
		newTimers[i-1] += timers[i]
	}
	return newTimers
}
