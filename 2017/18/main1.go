package main

import (
	"fmt"
	"os"
)

func main() {
	NewVM().Run()
}

type VM struct {
	Registers          map[string]int
	Cmds               []Cmd
	Pos                int
	LastSoundFrequency int
}

func NewVM() *VM {
	cmds := parse()
	return &VM{
		Registers: make(map[string]int),
		Cmds:      cmds,
		Pos:       0,
	}
}

func (vm *VM) Run() {
	for {
		vm.Exec()

		if vm.Pos < 0 || vm.Pos >= len(vm.Cmds) {
			return
		}
	}
}

func (vm *VM) Exec() {
	cmd := vm.Cmds[vm.Pos]

	switch cmd.Op {
	case "snd":
		vm.LastSoundFrequency = vm.getVal(cmd.X)
	case "set":
		vm.Registers[cmd.X] = vm.getVal(cmd.Y)
	case "add":
		vm.Registers[cmd.X] += vm.getVal(cmd.Y)
	case "mul":
		vm.Registers[cmd.X] *= vm.getVal(cmd.Y)
	case "mod":
		vm.Registers[cmd.X] %= vm.getVal(cmd.Y)
	case "rcv":
		if vm.getVal(cmd.X) != 0 {
			fmt.Println(vm.LastSoundFrequency)
			os.Exit(0)
		}
	case "jgz":
		if vm.getVal(cmd.X) > 0 {
			vm.Pos += vm.getVal(cmd.Y)
			return
		}
	}
	vm.Pos++
}
