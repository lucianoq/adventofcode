package main

import (
	"fmt"
	"time"
)

func main() {
	cmds := parse()
	out1 := make(chan int, 100000)
	out2 := make(chan int, 100000)
	p0 := NewVM(0, cmds, out1, out2)
	p1 := NewVM(1, cmds, out2, out1)

	go p0.Run()
	go p1.Run()

	time.Sleep(200 * time.Millisecond)
	fmt.Println(p1.NumSend)
}

type VM struct {
	ID        int
	Registers map[string]int
	Cmds      []Cmd
	Pos       int
	In        <-chan int
	Out       chan<- int
	NumSend   int
}

func NewVM(id int, cmds []Cmd, out chan<- int, in <-chan int) *VM {
	return &VM{
		ID:        id,
		Registers: map[string]int{"p": id},
		Cmds:      cmds,
		In:        in,
		Out:       out,
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
		vm.NumSend++
		vm.Out <- vm.getVal(cmd.X)
	case "set":
		vm.Registers[cmd.X] = vm.getVal(cmd.Y)
	case "add":
		vm.Registers[cmd.X] += vm.getVal(cmd.Y)
	case "mul":
		vm.Registers[cmd.X] *= vm.getVal(cmd.Y)
	case "mod":
		vm.Registers[cmd.X] %= vm.getVal(cmd.Y)
	case "rcv":
		vm.Registers[cmd.X] = <-vm.In
	case "jgz":
		if vm.getVal(cmd.X) > 0 {
			vm.Pos += vm.getVal(cmd.Y)
			return
		}
	}
	vm.Pos++
}
