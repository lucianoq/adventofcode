package main

import "fmt"

func main() {
	lines := parseInput()
	ops, lines := lines[len(lines)-1], lines[:len(lines)-1]

	grandTotal := 0
	for i := 0; i < len(ops); i++ {
		if ops[i] == '*' || ops[i] == '+' {
			grandTotal += applyHorizontal(lines, i, ops[i])
		}
	}
	fmt.Println(grandTotal)
}

func applyHorizontal(lines []string, idx int, op byte) int {
	result := neutral[op]
	for _, line := range lines {
		num := readHorizontal(line, idx)
		result = fn[op](result, num)
	}
	return result
}

func readHorizontal(line string, idx int) int {
	num := 0
	for line[idx] == ' ' {
		idx++
	}
	for i := idx; i < len(line); i++ {
		if line[i] < '0' || line[i] > '9' {
			break
		}
		num = 10*num + int(line[i]-'0')
	}
	return num
}
