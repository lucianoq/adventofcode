package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cmd struct {
	OpCode string
	R      string
	Offset int
}

func parse() []Cmd {
	program := make([]Cmd, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		items := strings.Split(line, " ")
		switch items[0] {
		case "hlf", "tpl", "inc":
			program = append(program, Cmd{items[0], items[1], 0})
		case "jmp":
			offset, err := strconv.Atoi(items[1])
			if err != nil {
				log.Fatal(err)
			}
			program = append(program, Cmd{"jmp", "", offset})
		case "jie", "jio":
			r := strings.TrimRight(items[1], ",")
			offset, err := strconv.Atoi(items[2])
			if err != nil {
				log.Fatal(err)
			}
			program = append(program, Cmd{items[0], r, offset})
		default:
			log.Fatal("wrong opcode")
		}
	}

	return program
}

func run(program []Cmd, reg map[string]uint64) uint64 {
	i := 0
	for {
		if i < 0 || i >= len(program) {
			break
		}

		c := program[i]

		log.Printf("[a=%d b=%d]  i=%d -> %v", reg["a"], reg["b"], i, c)

		switch c.OpCode {
		case "hlf":
			reg[c.R] /= 2
		case "tpl":
			reg[c.R] *= 3
		case "inc":
			reg[c.R]++
		case "jmp":
			i += c.Offset
			continue
		case "jie":
			if reg[c.R]%2 == 0 {
				i += c.Offset
				continue
			}
		case "jio":
			if reg[c.R] == 1 {
				i += c.Offset
				continue
			}
		}
		i++
	}

	return reg["b"]
}
