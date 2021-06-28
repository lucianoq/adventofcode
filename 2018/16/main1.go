package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	tests := read1(os.Stdin)
	mainCount := 0
	for _, t := range tests {
		count := 0
		for _, f := range opFuncs {
			register = t.Before
			f(t.In.A, t.In.B, t.In.C)
			if register == t.After {
				count++
			}
		}

		if count >= 3 {
			mainCount++
		}
	}
	fmt.Println(mainCount)
}

type Test struct {
	Before [4]int
	In     Instruction
	After  [4]int
}

func read1(reader io.Reader) []Test {
	tests := make([]Test, 0)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

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
	return tests
}
