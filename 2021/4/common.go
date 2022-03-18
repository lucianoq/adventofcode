package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/lucianoq/container/set"
)

type Board [5][5]int

func (b *Board) Wins(drawn set.Set[int]) bool {
Row:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !drawn.Contains(b[i][j]) {
				continue Row
			}
		}
		return true
	}

Column:
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			if !drawn.Contains(b[i][j]) {
				continue Column
			}
		}
		return true
	}

	return false
}

func (b *Board) SumUnmarked(drawn set.Set[int]) int {
	sumUnmarked := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !drawn.Contains(b[i][j]) {
				sumUnmarked += b[i][j]
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
			board[i][j] = num
		}
	}
	return board
}
