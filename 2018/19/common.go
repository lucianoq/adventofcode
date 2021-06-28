package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func read(reader io.Reader) []Instruction {
	instructions := make([]Instruction, 0)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			_, _ = fmt.Sscanf(line, "#ip %d", &ipRegister)
			continue
		}

		var o string
		var a, b, c int
		_, _ = fmt.Sscanf(line, "%s %d %d %d", &o, &a, &b, &c)

		instructions = append(instructions, Instruction{
			FName: o,
			F:     opFuncs[o],
			A:     a,
			B:     b,
			C:     c,
		})
	}
	return instructions
}
