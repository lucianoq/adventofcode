package main

import (
	"bufio"
	"os"
)

const Size = 50

type P struct {
	X, Y int
}

func (p P) Valid() bool {
	return p.X >= 0 && p.X < Size && p.Y >= 0 && p.Y < Size
}

func parseInput() map[byte][]P {
	scanner := bufio.NewScanner(os.Stdin)

	freqs := map[byte][]P{}
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		for j, c := range line {
			if c == '.' {
				continue
			}

			ch := byte(c)

			freqs[ch] = append(freqs[ch], P{i, j})
		}
	}

	return freqs
}
