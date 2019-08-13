package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var reg = make(map[string]int)

func run() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		ff := strings.Fields(line)
		switch ff[0] {
		case "cpy":
			x := getVal(ff[1])
			reg[ff[2]] = x
		case "inc":
			reg[ff[1]]++
		case "dec":
			reg[ff[1]]--
		case "jnz":
			x := getVal(ff[1])
			if x != 0 {
				n, _ := strconv.Atoi(ff[2])
				i += n - 1
			}
		}
	}
}

func getVal(x string) int {
	switch x {
	case "a", "b", "c", "d":
		return reg[x]
	default:
		n, _ := strconv.Atoi(x)
		return n
	}
}
