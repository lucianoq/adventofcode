package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Board [5][5]int

func (b *Board) Wins(drawn map[int]struct{}) bool {
Row:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			_, marked := drawn[b[i][j]]
			if !marked {
				continue Row
			}
		}
		return true
	}

Column:
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			_, marked := drawn[b[i][j]]
			if !marked {
				continue Column
			}
		}
		return true
	}

	return false
}

func (b *Board) SumUnmarked(drawn map[int]struct{}) int {
	sumUnmarked := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			_, marked := drawn[b[i][j]]
			if !marked {
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
