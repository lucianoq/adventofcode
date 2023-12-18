package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Op struct {
	Dir   string
	Steps int
	Color string
}

type P struct{ R, C int }

var Delta = map[string]P{
	"U": {-1, 0},
	"R": {0, 1},
	"D": {1, 0},
	"L": {0, -1},
}

func parse() []Op {
	scanner := bufio.NewScanner(os.Stdin)
	var ops []Op
	for scanner.Scan() {
		line := scanner.Text()

		var op Op
		_, _ = fmt.Sscanf(line, "%s %d %s", &op.Dir, &op.Steps, &op.Color)
		op.Color = strings.Trim(op.Color, "()#")
		ops = append(ops, op)
	}
	return ops
}
