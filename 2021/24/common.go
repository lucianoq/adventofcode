package main

import (
	"fmt"

	"github.com/lucianoq/container/set"
)

type Status struct{ DigitIdx, Z int }

var uselessBranches = set.New[Status]()

func main() {
	model := recursive([]int{}, 0)
	fmt.Println(model)
}

func toInt(list []int) int {
	res := 0
	for _, n := range list {
		res = 10*res + n
	}
	return res
}

func alu(z, digit, digitIdx int) int {
	x := z%26 + input[digitIdx].AddToZmod26
	z /= input[digitIdx].DivideZBy
	if digit == x {
		x = 0
	} else {
		x = 1
	}
	return z*(x*25+1) + x*(digit+input[digitIdx].AddToW)
}

var input = [14]struct {
	AddToZmod26 int
	DivideZBy   int
	AddToW      int
}{
	{10, 1, 2},
	{10, 1, 4},
	{14, 1, 8},
	{11, 1, 7},
	{14, 1, 12},
	{-14, 26, 7},
	{0, 26, 10},
	{10, 1, 14},
	{-10, 26, 2},
	{13, 1, 6},
	{-12, 26, 8},
	{-3, 26, 11},
	{-11, 26, 5},
	{-2, 26, 11},
}
