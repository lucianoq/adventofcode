package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	tests := read1()
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

func Dump(tests []Test) {
	for _, t := range tests {
		fmt.Printf("B [%d %d %d %d]\n", t.Before[0], t.Before[1], t.Before[2], t.Before[3])
		fmt.Printf("%d %d %d %d\n", t.In.OpCode, t.In.A, t.In.B, t.In.C)
		fmt.Printf("A [%d %d %d %d]\n", t.After[0], t.After[1], t.After[2], t.After[3])
	}
}

func read1() []Test {
	tests := make([]Test, 0)

	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)

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
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return tests
}
