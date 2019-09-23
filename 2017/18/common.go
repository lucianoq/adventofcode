package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type OpCode int

type Cmd struct {
	Op string
	X  string
	Y  string
}

func parse() []Cmd {
	cmds := make([]Cmd, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		ff := strings.Fields(line)

		cmd := Cmd{
			Op: ff[0],
			X:  ff[1],
		}

		if len(ff) > 2 {
			cmd.Y = ff[2]
		}
		cmds = append(cmds, cmd)
	}

	return cmds
}

func (vm *VM) getVal(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		return vm.Registers[s]
	}
	return num
}
