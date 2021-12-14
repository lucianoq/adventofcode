package main

import (
	"bufio"
	"os"
	"strings"
)

func parse() (string, map[string]string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	start := scanner.Text()

	rules := map[string]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		ff := strings.Split(line, " -> ")
		rules[ff[0]] = ff[1]
	}

	return start, rules
}

func minMax(count map[uint8]int) (int, int) {
	min, max := 1<<63-1, -1<<63
	for _, v := range count {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}
