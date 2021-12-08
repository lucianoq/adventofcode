package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parse() []int {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line := scanner.Text()

	ff := strings.Split(strings.TrimSpace(line), ",")

	list := []int{}
	for _, f := range ff {
		num, _ := strconv.Atoi(f)

		list = append(list, num)
	}
	return list
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func minMax(list []int) (int, int) {
	min, max := 1<<63-1, -1<<63
	for _, x := range list {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	return min, max
}
