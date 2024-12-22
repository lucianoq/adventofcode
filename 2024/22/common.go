package main

import (
	"bufio"
	"os"
	"strconv"
)

func parseInput() []int {
	var list []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		list = append(list, num)
	}
	return list
}

func compute(secret int) int {
	secret = (secret ^ (secret << 6)) % 16777216
	secret = (secret ^ (secret >> 5)) % 16777216
	secret = (secret ^ (secret << 11)) % 16777216
	return secret
}
