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

func evolve(timers map[int]int) {
	for i := 0; i <= 8; i++ {
		timers[i-1] = timers[i]
	}
	timers[8] = timers[-1]
	timers[6] += timers[-1]
	timers[-1] = 0
}
