package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	Indicator  uint16
	Buttons    []uint16
	ButtonsInt [][]int
	Joltage    []int
}

func parseInput() []Machine {
	var list []Machine

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		tokens := strings.Split(scanner.Text(), " ")

		indicatorS := strings.Trim(tokens[0], "[]")
		var indicator uint16
		for i, c := range indicatorS {
			if c == '#' {
				indicator |= 1 << uint16(len(indicatorS)-i-1)
			}
		}

		var buttons []uint16
		var buttonsInt [][]int
		for i := 1; i < len(tokens)-1; i++ {
			buttonS := strings.Split(strings.Trim(tokens[i], "()"), ",")
			buttonInt := make([]int, len(buttonS))
			var button uint16
			for j, b := range buttonS {
				n, _ := strconv.Atoi(b)
				buttonInt[j] = n
				button |= 1 << uint16(len(indicatorS)-n-1)
			}
			buttons = append(buttons, button)
			buttonsInt = append(buttonsInt, buttonInt)
		}

		joltageS := strings.Split(strings.Trim(tokens[len(tokens)-1], "{}"), ",")
		joltage := make([]int, len(joltageS))
		for i := range joltageS {
			joltage[i], _ = strconv.Atoi(joltageS[i])
		}
		list = append(list, Machine{indicator, buttons, buttonsInt, joltage})
	}
	return list
}
