package main

import "fmt"

func main() {
	lines := parseInput()
	ops, lines := lines[len(lines)-1], lines[:len(lines)-1]

	grandTotal := 0
	for i := 0; i < len(ops); i++ {
		if ops[i] == '*' || ops[i] == '+' {
			var j int
			for j = i + 1; j < len(ops); j++ {
				if ops[j] == '*' || ops[j] == '+' {
					j--
					break
				}
			}
			grandTotal += applyVertical(lines, i, j, ops[i])
		}
	}

	fmt.Println(grandTotal)
}

func applyVertical(lines []string, i, j int, op byte) int {
	result := neutral[op]
	for k := i; k < j; k++ {
		num := readVertical(lines, k)
		result = fn[op](result, num)
	}
	return result
}

func readVertical(lines []string, idx int) int {
	num := 0
	for _, line := range lines {
		if line[idx] == ' ' {
			continue
		}
		num = num*10 + int(line[idx]-'0')
	}
	return num
}
