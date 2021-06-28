package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	tests, instructions := read(os.Stdin)

	for _, t := range tests {
		for idx, f := range opFuncs {
			register = t.Before
			f(t.In.A, t.In.B, t.In.C)
			if register == t.After {
				candidates[t.In.OpCode][idx] = true
			}
		}
	}

	Solve()

	register = [4]int{0, 0, 0, 0}
	for _, i := range instructions {
		solved[i.OpCode](i.A, i.B, i.C)
	}
	fmt.Println(register[0])
}

func Solve() {
	for len(solved) < 16 {
		for opCode, c := range candidates {
			keys := make([]int, 0)
			for k := range c {
				keys = append(keys, k)
			}
			if len(keys) == 1 {
				solved[opCode] = opFuncs[keys[0]]
				for _, c2 := range candidates {
					delete(c2, keys[0])
				}
			}
		}
	}
}

type Test struct {
	Before [4]int
	In     Instruction
	After  [4]int
}

func read(reader io.Reader) ([]Test, []Instruction) {
	tests := make([]Test, 0)

	scanner := bufio.NewScanner(reader)

	emptyLines := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			emptyLines++
			if emptyLines == 2 {
				break
			}
			continue
		}
		emptyLines = 0
		if strings.HasPrefix(line, "Before") {
			var t Test

			_, _ = fmt.Sscanf(line, "Before: [%d, %d, %d, %d]", &t.Before[0], &t.Before[1], &t.Before[2], &t.Before[3])

			scanner.Scan()
			line = strings.TrimSpace(scanner.Text())
			var o, a, b, c int
			_, _ = fmt.Sscanf(line, "%d %d %d %d", &o, &a, &b, &c)
			t.In.OpCode, t.In.A, t.In.B, t.In.C = o, a, b, c

			scanner.Scan()
			line = strings.TrimSpace(scanner.Text())
			_, _ = fmt.Sscanf(line, "After: [%d, %d, %d, %d]", &t.After[0], &t.After[1], &t.After[2], &t.After[3])

			tests = append(tests, t)
		}
	}

	instructions := make([]Instruction, 0)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		var o, a, b, c int
		_, _ = fmt.Sscanf(line, "%d %d %d %d", &o, &a, &b, &c)

		instructions = append(instructions, Instruction{
			OpCode: o,
			A:      a,
			B:      b,
			C:      c,
		})
	}

	return tests, instructions
}
