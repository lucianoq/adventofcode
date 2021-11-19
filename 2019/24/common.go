package main

import (
	"bufio"
	"os"
)

const (
	Side   = 5
	Square = Side * Side
)

func parse() [Square]bool {
	res := [Square]bool{}

	s := bufio.NewScanner(os.Stdin)
	for row := 0; s.Scan(); row++ {
		line := s.Text()
		for col := 0; col < Side; col++ {
			if line[col] == '#' {
				res[row*Side+col] = true
			} else {
				res[row*Side+col] = false
			}
		}
	}
	return res
}
