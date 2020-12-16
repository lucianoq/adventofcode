package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func game(input []int, nth int) int {
	spoken := map[int]int{}

	var (
		toSpeak int
		age     int
	)

	// returns -1 if first occurrence of num
	// returns its age otherwise
	isFirstOrAge := func(num, turn int) int {
		t, ok := spoken[num]
		if ok {
			return turn - t
		} else {
			return -1
		}
	}

	// from 1 to len(input) inject input
	for turn := 1; turn <= len(input); turn++ {
		toSpeak = input[turn-1]
		age = isFirstOrAge(toSpeak, turn)
		spoken[toSpeak] = turn
	}

	for turn := len(input) + 1; turn <= nth; turn++ {
		if age == -1 {
			toSpeak = 0
		} else {
			toSpeak = age
		}

		age = isFirstOrAge(toSpeak, turn)
		spoken[toSpeak] = turn
	}

	return toSpeak
}

func parse() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ff := strings.Split(strings.TrimSpace(scanner.Text()), ",")

	var input []int
	for _, f := range ff {
		num, _ := strconv.Atoi(f)
		input = append(input, num)
	}
	return input
}
