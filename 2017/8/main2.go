package main

import (
	"fmt"
	"math"
)

func main() {
	program := parse()
	memory := make(map[string]int)
	max := math.MinInt64

	for _, i := range program {

		// run instruction
		val := runVal(memory, i)

		if val > max {
			max = val
		}
	}

	fmt.Println(max)
}

func runVal(memory map[string]int, i Instruction) int {
	switch {
	case i.TestOp == ">" && memory[i.TestReg] > i.TestVal,
		i.TestOp == ">=" && memory[i.TestReg] >= i.TestVal,
		i.TestOp == "<" && memory[i.TestReg] < i.TestVal,
		i.TestOp == "<=" && memory[i.TestReg] <= i.TestVal,
		i.TestOp == "==" && memory[i.TestReg] == i.TestVal,
		i.TestOp == "!=" && memory[i.TestReg] != i.TestVal:
		memory[i.ResultReg] += i.Increment
	}
	return memory[i.ResultReg]
}
