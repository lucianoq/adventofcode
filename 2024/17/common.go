package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseInput() (int, int, int, []int) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	a, _ := strconv.Atoi(strings.Fields(scanner.Text())[2])

	scanner.Scan()
	b, _ := strconv.Atoi(strings.Fields(scanner.Text())[2])

	scanner.Scan()
	c, _ := strconv.Atoi(strings.Fields(scanner.Text())[2])

	scanner.Scan()

	scanner.Scan()
	var code []int
	for _, s := range strings.Split(strings.Fields(scanner.Text())[1], ",") {
		n, _ := strconv.Atoi(s)
		code = append(code, n)
	}
	return a, b, c, code
}

func run(a, b, c int, code []int) []int {
	output := make([]int, 0)

	for ip := 0; ip < len(code); {
		var (
			op    = code[ip]
			arg   = code[ip+1]
			combo = 0
		)

		switch arg {
		case 1, 2, 3:
			combo = arg
		case 4:
			combo = a
		case 5:
			combo = b
		case 6:
			combo = c
		}

		switch op {
		case 0: // adv
			a >>= combo
		case 1: // bxl
			b ^= arg
		case 2: // bst
			b = combo % 8
		case 3: // jnz
			if a != 0 {
				ip = arg
				continue
			}
		case 4: // bxc
			b ^= c
		case 5: // out
			output = append(output, combo%8)
		case 6: // bdv
			b = a >> combo
		case 7: // cdv
			c = a >> combo
		}

		ip += 2
	}

	return output
}
