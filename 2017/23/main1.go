package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cmd struct {
	Op, X, Y string
}

type VM struct {
	Registers map[string]int
	Cmds      []Cmd
	Pos       int
	NumMul    int
}

func main() {
	p := &VM{
		Registers: map[string]int{},
		Cmds:      parse(),
	}

	p.Run()

	fmt.Println(p.NumMul)
}

func (vm *VM) Run() {
	for i := 0; ; i++ {
		if i&0xffffff == 0 {
			log.Println(i)
		}
		vm.Exec()

		if vm.Pos < 0 || vm.Pos >= len(vm.Cmds) {
			return
		}
	}
}

func (vm *VM) Exec() {
	cmd := vm.Cmds[vm.Pos]

	switch cmd.Op {
	case "set":
		vm.Registers[cmd.X] = vm.getVal(cmd.Y)
	case "sub":
		vm.Registers[cmd.X] -= vm.getVal(cmd.Y)
	case "mul":
		vm.NumMul++
		vm.Registers[cmd.X] *= vm.getVal(cmd.Y)
	case "jnz":
		if vm.getVal(cmd.X) != 0 {
			vm.Pos += vm.getVal(cmd.Y)
			return
		}
	}
	vm.Pos++
}

func (vm *VM) getVal(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		return vm.Registers[s]
	}
	return num
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
