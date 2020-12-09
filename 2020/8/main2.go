package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Op struct {
	Code string
	Val  int
}

type VM struct {
	cmds        []Op
	cursor      int
	Accumulator int
	visited     map[int]struct{}
}

func NewVM(cmds []Op) *VM {
	return &VM{
		cmds:        cmds,
		cursor:      0,
		Accumulator: 0,
		visited:     map[int]struct{}{},
	}
}

func (v *VM) Run() error {
	for v.cursor >= 0 && v.cursor < len(v.cmds) {
		if _, ok := v.visited[v.cursor]; ok {
			// found loop
			return errors.New("loop")
		}
		v.visited[v.cursor] = struct{}{}

		curr := v.cmds[v.cursor]

		switch curr.Code {
		case "acc":
			v.Accumulator += curr.Val
			v.cursor += 1
		case "jmp":
			v.cursor += curr.Val
		case "nop":
			v.cursor += 1
		}
	}

	// program terminates
	return nil
}

func parse() []Op {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := []Op{}

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		val, _ := strconv.Atoi(arr[1])

		cmds = append(cmds, Op{
			Code: arr[0],
			Val:  val,
		})
	}
	return cmds
}

func main() {
	cmds := parse()

	for i := 0; i < len(cmds); i++ {

		if cmds[i].Code == "acc" {
			continue
		}

		// create a copy
		newCmds := make([]Op, len(cmds))
		copy(newCmds, cmds)

		// swap
		switch cmds[i].Code {
		case "jmp":
			newCmds[i].Code = "nop"
		case "nop":
			newCmds[i].Code = "jmp"
		}

		vm := NewVM(newCmds)
		err := vm.Run()
		if err == nil {
			fmt.Println(vm.Accumulator)
			return
		}
	}
}
