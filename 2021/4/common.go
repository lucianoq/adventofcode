package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	grid   [5][5]int
	marked [5][5]bool
}

func (b *Board) Mark(num int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.grid[i][j] == num {
				b.marked[i][j] = true
			}
		}
	}
}

func (b *Board) Wins() bool {
Row:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.marked[i][j] {
				continue Row
			}
		}
		return true
	}

Column:
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			if !b.marked[i][j] {
				continue Column
			}
		}
		return true
	}

	return false
}

func (b *Board) SumUnmarked() int {
	sumUnmarked := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.marked[i][j] {
				sumUnmarked += b.grid[i][j]
			}
		}
	}
	return sumUnmarked
}

func parse() ([]int, []*Board) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	ff := strings.Split(line, ",")
	draws := make([]int, 0, len(ff))
	for _, f := range ff {
		num, _ := strconv.Atoi(f)
		draws = append(draws, num)
	}

	boards := make([]*Board, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			boards = append(boards, readBoard(scanner))
		}
	}

	return draws, boards
}

func readBoard(scanner *bufio.Scanner) *Board {
	board := &Board{}

	for i := 0; i < 5; i++ {
		scanner.Scan()
		line := strings.ReplaceAll(scanner.Text(), "  ", " ")
		ff := strings.Split(strings.TrimSpace(line), " ")
		for j := 0; j < 5; j++ {
			num, _ := strconv.Atoi(ff[j])
			board.grid[i][j] = num
		}
	}
	return board
}
