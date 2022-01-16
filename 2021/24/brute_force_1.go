package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	alu := NewALU(parse())

Model:
	for model := 99999999999999; model >= 11111111111111; model-- {

		s := strconv.Itoa(model)
		for _, c := range s {
			if c == '0' {
				continue Model
			}
		}

		if model%107777 == 0 {
			log.Println(model)
		}

		if alu.Valid(model) {
			fmt.Println(model)
			return
		}
		alu.Reset()
	}
}

func NewALU(instructions []Instruction) *ALU {
	return &ALU{
		Memory: map[string]int{
			"w": 0,
			"x": 0,
			"y": 0,
			"z": 0,
		},
		Program: instructions,
	}
}

type Instruction struct {
	Cmd        string
	Arg1, Arg2 string
}

func parse() []Instruction {
	scanner := bufio.NewScanner(os.Stdin)

	instructions := []Instruction{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		ff := strings.Split(line, " ")

		inst := Instruction{
			Cmd:  ff[0],
			Arg1: ff[1],
		}
		if len(ff) > 2 {
			inst.Arg2 = ff[2]
		}

		instructions = append(instructions, inst)
	}
	return instructions
}

type ALU struct {
	Memory  map[string]int
	Program []Instruction
}

func (a *ALU) Reset() {
	for k := range a.Memory {
		a.Memory[k] = 0
	}
}

func (a *ALU) Valid(num int) bool {
	inputIdx, input := 0, strconv.Itoa(num)

	for _, i := range a.Program {
		switch i.Cmd {
		case "inp":
			a.Memory[i.Arg1] = int(input[inputIdx] - '0')
			inputIdx++
		case "add":
			a.Memory[i.Arg1] += a.getVal(i.Arg2)
		case "mul":
			a.Memory[i.Arg1] *= a.getVal(i.Arg2)
		case "div":
			a.Memory[i.Arg1] /= a.getVal(i.Arg2)
		case "mod":
			a.Memory[i.Arg1] %= a.getVal(i.Arg2)
		case "eql":
			if a.Memory[i.Arg1] == a.getVal(i.Arg2) {
				a.Memory[i.Arg1] = 1
			} else {
				a.Memory[i.Arg1] = 0
			}
		}
	}

	return a.Memory["z"] == 0
}

func (a *ALU) getVal(x string) int {
	switch x {
	case "w", "x", "y", "z":
		return a.Memory[x]
	}
	val, _ := strconv.Atoi(x)
	return val
}
