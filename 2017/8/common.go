package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	TestReg   string
	TestOp    string
	TestVal   int
	Increment int
	ResultReg string
}

func parse() []Instruction {
	scanner := bufio.NewScanner(os.Stdin)

	program := make([]Instruction, 0)

	for scanner.Scan() {
		line := scanner.Text()

		ff := strings.Fields(line)

		increment, _ := strconv.Atoi(ff[2])
		if ff[1] == "dec" {
			increment *= -1
		}

		testValue, _ := strconv.Atoi(ff[6])

		program = append(program, Instruction{
			TestReg:   ff[4],
			TestOp:    ff[5],
			TestVal:   testValue,
			Increment: increment,
			ResultReg: ff[0],
		})
	}
	return program
}
