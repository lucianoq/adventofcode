package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	instructions := read()

	i, _ := strconv.Atoi(os.Args[1])

	register = [6]int{i, 0, 0, 0, 0, 0}
	ip = 0

	for {
		if ip < 0 || ip >= len(instructions) {
			log.Println("Halt")
			break
		}

		Dump(instructions[ip])
	}

	fmt.Println(register[0])
	time.Sleep(2 * time.Second)

}

func Dump(ins Instruction) {
	fmt.Printf("ip=%d ", ip)
	dumpReg()
	fmt.Printf(" %s %d %d %d ", ins.FName, ins.A, ins.B, ins.C)
	ins.Exec()
	dumpReg()
	fmt.Println()
}

func dumpReg() {
	fmt.Printf("[%d", register[0])
	for i := 1; i < 6; i++ {
		fmt.Printf(", %d", register[i])
	}
	fmt.Print("]")
}

func read() []Instruction {
	instructions := make([]Instruction, 0)

	f, _ := os.Open("input")
	scanner := bufio.NewScanner(f)

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
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return instructions
}
