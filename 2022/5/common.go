package main

import (
	"bufio"
	"fmt"
	"os"
)

type Move struct{ Qty, From, To int }

func parse() (map[int][]byte, []Move) {
	scanner := bufio.NewScanner(os.Stdin)

	stacks := map[int][]byte{}
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		for i := 1; i < len(line); i += 4 {
			if line[i] >= 'A' && line[i] <= 'Z' {
				stackID := (i / 4) + 1
				stacks[stackID] = append(stacks[stackID], line[i])
			}
		}
	}

	// reverse lists
	for _, stack := range stacks {
		for i := 0; i < len(stack)/2; i++ {
			stack[i], stack[len(stack)-1-i] = stack[len(stack)-1-i], stack[i]
		}
	}

	moves := []Move{}
	for scanner.Scan() {
		line := scanner.Text()
		var qty, from, to int
		fmt.Sscanf(line, "move %d from %d to %d\n", &qty, &from, &to)
		moves = append(moves, Move{qty, from, to})
	}

	return stacks, moves
}

func printTop(stacks map[int][]byte) {
	s := ""
	for i := 1; i <= len(stacks); i++ {
		last := len(stacks[i]) - 1
		s += string(stacks[i][last])
	}
	fmt.Println(s)
}
