package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Pins   = 5
	Height = 7
)

type Schematic [Pins]int

func main() {
	keys, locks := parseInput()

	count := 0
	for _, key := range keys {
		for _, lock := range locks {
			if !overlap(key, lock) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func overlap(key, lock Schematic) bool {
	for i := 0; i < Pins; i++ {
		if key[i]+lock[i] > Height-2 {
			return true
		}
	}
	return false
}

func parseInput() ([]Schematic, []Schematic) {
	var keys, locks []Schematic

	scanner := bufio.NewScanner(os.Stdin)
	for {
		m := [Height]string{}
		for i := 0; i < Height; i++ {
			if !scanner.Scan() {
				return keys, locks
			}
			m[i] = scanner.Text()
		}

		val, isLock := parseSchema(m)
		if isLock {
			locks = append(locks, val)
		} else {
			keys = append(keys, val)
		}
		scanner.Scan()
	}
}

func parseSchema(m [Height]string) (Schematic, bool) {
	res := Schematic{}
	for j := 0; j < Pins; j++ {
		count := 0
		// ignore first and last line
		for i := 1; i < Height-1; i++ {
			if m[i][j] == '#' {
				count++
			}
		}
		res[j] = count
	}
	return res, m[0] == "#####" && m[Height-1] == "....."
}
