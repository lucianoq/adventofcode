package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parse() [][]int {
	scanner := bufio.NewScanner(os.Stdin)
	list := [][]int{}
	for scanner.Scan() {
		sequence := []int{}
		for _, f := range strings.Fields(scanner.Text()) {
			n, _ := strconv.Atoi(f)
			sequence = append(sequence, n)
		}
		list = append(list, sequence)
	}
	return list
}

func buildPascal(firstSeq []int) [][]int {
	p := [][]int{firstSeq}
	for {
		last := len(p) - 1

		done := true
		newRow := make([]int, len(p[last])-1)
		for i := 0; i < len(p[last])-1; i++ {
			newRow[i] = p[last][i+1] - p[last][i]
			if newRow[i] != 0 {
				done = false
			}
		}
		p = append(p, newRow)

		if done {
			return p
		}
	}
}
