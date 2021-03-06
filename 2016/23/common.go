package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	cpy uint8 = iota
	inc
	dec
	jnz
	tgl
)

type Cmd struct {
	Op   uint8
	X, Y string
}

var (
	reg    [4]int
	cmds   []Cmd
	cursor int
)

func (c Cmd) Run() {
	//log.Printf("{i=%d} reg[%d %d %d %d] \t Running: %s (%s %s)\t", cursor, reg[0], reg[1], reg[2], reg[2], c.Op, c.X, c.Y)

	switch c.Op {
	case cpy:
		x := getVal(c.X)
		if r, ok := isReg(c.Y); ok {
			reg[r] = x
		}
	case inc:
		if r, ok := isReg(c.X); ok {
			reg[r]++
		}
	case dec:
		if r, ok := isReg(c.X); ok {
			reg[r]--
		}

	case jnz:
		x := getVal(c.X)
		if x != 0 {
			n := getVal(c.Y)
			cursor += n - 1
		}
	case tgl:
		x := getVal(c.X)
		target := cursor + x

		if target <= 0 || target >= len(cmds) {
			return
		}

		cmd := cmds[target]

		switch cmd.Op {
		case cpy:
			cmd.Op = jnz

		case inc:
			cmd.Op = dec

		case dec:
			cmd.Op = inc

		case jnz:
			cmd.Op = cpy

		case tgl:
			cmd.Op = inc
		}

		cmds[target] = cmd
	}
}

func parse() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		ff := strings.Fields(line)

		var cmd = Cmd{
			X: ff[1],
		}

		switch ff[0] {
		case "cpy":
			cmd.Op = cpy
		case "inc":
			cmd.Op = inc
		case "dec":
			cmd.Op = dec
		case "jnz":
			cmd.Op = jnz
		case "tgl":
			cmd.Op = tgl
		}

		if ff[0] == "cpy" || ff[0] == "jnz" {
			cmd.Y = ff[2]
		}

		cmds = append(cmds, cmd)
	}
}

func getVal(x string) int {
	switch x {
	case "a", "b", "c", "d":
		r := []byte(x)[0]
		return reg[r-'a']
	default:
		n, _ := strconv.Atoi(x)
		return n
	}
}

func isReg(s string) (int, bool) {
	switch s {
	case "a", "b", "c", "d":
		r := []byte(s)
		return int(r[0] - 'a'), true
	}
	return 0, false
}
