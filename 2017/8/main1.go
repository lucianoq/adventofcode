package main

import (
	"fmt"
	"math"
)

func main() {
	program := parse()
	memory := make(map[string]int)

	for _, i := range program {
		// run instruction
		run(memory, i)
	}

	fmt.Println(max(memory))
}

func run(memory map[string]int, i Instruction) {
	switch {
	case i.TestOp == ">" && memory[i.TestReg] > i.TestVal,
		i.TestOp == ">=" && memory[i.TestReg] >= i.TestVal,
		i.TestOp == "<" && memory[i.TestReg] < i.TestVal,
		i.TestOp == "<=" && memory[i.TestReg] <= i.TestVal,
		i.TestOp == "==" && memory[i.TestReg] == i.TestVal,
		i.TestOp == "!=" && memory[i.TestReg] != i.TestVal:
		memory[i.ResultReg] += i.Increment
	}
}

func max(memory map[string]int) int {
	max := math.MinInt64
	for _, v := range memory {
		if v > max {
			max = v
		}
	}
	return max
}
