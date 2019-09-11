package main

import (
	"bufio"
	"os"
	"strconv"
)

func parse() []int {
	scanner := bufio.NewScanner(os.Stdin)
	list := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)
		list = append(list, n)
	}
	return list
}
