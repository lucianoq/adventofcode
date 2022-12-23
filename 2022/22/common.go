package main

import (
	"bufio"
	"fmt"
	"os"
)

type P struct{ X, Y int }

type Dir int8

const (
	N Dir = iota
	E
	S
	W
)

func (dir Dir) Rotate(direction rune) Dir {
	switch direction {
	case 'R':
		return (dir + 1) % 4
	case 'L':
		return (dir - 1 + 4) % 4
	}
	return dir
}

func (dir Dir) Points() int {
	return (int(dir) + 3) % 4
}

type Movement struct {
	Steps  int
	Rotate rune
}

var (
	Map       = map[P]bool{}
	Size      int
	Movements []Movement
	Dirs      = [4]P{
		{-1, 0}, // N
		{0, 1},  // E
		{1, 0},  // S
		{0, -1}, // W
	}
)

type Human struct {
	Curr   P
	Facing Dir
}

func main() {
	parse()

	human := Human{
		Curr:   P{0, Size},
		Facing: E,
	}

	for _, mov := range Movements {
		human.Facing = human.Facing.Rotate(mov.Rotate)
		for i := 0; i < mov.Steps; i++ {
			if moved := human.Walk(); !moved {
				break
			}
		}
	}

	fmt.Println(1000*(human.Curr.X+1) + 4*(human.Curr.Y+1) + human.Facing.Points())
}

func parse() {
	scanner := bufio.NewScanner(os.Stdin)
	for r := 0; scanner.Scan(); r++ {
		line := scanner.Text()

		if line == "" {
			break
		}

		if r == 0 {
			Size = len(line) / 3
		}

		for c, char := range line {

			switch char {
			case ' ':
				continue
			case '#':
				Map[P{r, c}] = true
			case '.':
				Map[P{r, c}] = false
			}
		}
	}
	scanner.Scan()
	Movements = parsePath(scanner.Text())
}

func parsePath(path string) []Movement {
	movements := []Movement{}
	acc := 0
	for _, char := range path {

		switch char {

		case 'R':
			movements = append(movements, Movement{Steps: acc})
			acc = 0
			movements = append(movements, Movement{Rotate: 'R'})

		case 'L':
			movements = append(movements, Movement{Steps: acc})
			acc = 0
			movements = append(movements, Movement{Rotate: 'L'})

		default:
			acc = 10*acc + int(char-'0')

		}
	}
	movements = append(movements, Movement{Steps: acc})
	return movements
}
