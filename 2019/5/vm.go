package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cmd int

func (c Cmd) OpCode() int {
	// last 2 digits
	return int(c) % 100
}

func (c Cmd) Immediate(paramNum int) bool {
	// 10^2 =   100  for 1
	// 10^3 =  1000  for 2
	// 10^4 = 10000  for 3
	digit := int(math.Pow10(paramNum + 1))
	return int(c)/digit%10 == 1
}

type VM struct {
	Code []int // source code and memory, list of int
	Ip   int   // Instruction Pointer
}

func (v *VM) Load(r io.Reader) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	listStr := strings.Split(strings.TrimSpace(string(buf)), ",")
	listInt := make([]int, len(listStr))

	for i := range listStr {
		var err error
		listInt[i], err = strconv.Atoi(listStr[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	v.Code = listInt
	v.Ip = 0
}

func (v *VM) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		cmd := Cmd(v.Code[v.Ip])

		switch cmd.OpCode() {

		case 1:
			// add
			param1 := v.getParam(v.Ip+1, cmd.Immediate(1))
			param2 := v.getParam(v.Ip+2, cmd.Immediate(2))
			address := v.getParam(v.Ip+3, true)
			v.Code[address] = param1 + param2
			v.Ip += 4

		case 2:
			// multiply
			param1 := v.getParam(v.Ip+1, cmd.Immediate(1))
			param2 := v.getParam(v.Ip+2, cmd.Immediate(2))
			address := v.getParam(v.Ip+3, true)
			v.Code[address] = param1 * param2
			v.Ip += 4

		case 3:
			// read
			address := v.getParam(v.Ip+1, true)
			scanner.Scan()
			line := scanner.Text()
			line = strings.TrimSpace(line)
			input, _ := strconv.Atoi(line)
			v.Code[address] = input
			v.Ip += 2

		case 4:
			// write
			param1 := v.getParam(v.Ip+1, cmd.Immediate(1))
			fmt.Println(param1)
			v.Ip += 2

		case 5:
			// jump not zero
			param1 := v.getParam(v.Ip+1, cmd.Immediate(1))
			param2 := v.getParam(v.Ip+2, cmd.Immediate(2))
			if param1 != 0 {
				v.Ip = param2
			} else {
				v.Ip += 3
			}

		case 6:
			// jump zero
			param1 := v.getParam(v.Ip+1, cmd.Immediate(1))
			param2 := v.getParam(v.Ip+2, cmd.Immediate(2))
			if param1 == 0 {
				v.Ip = param2
			} else {
				v.Ip += 3
			}

		case 7:
			// less than
			param1 := v.getParam(v.Ip+1, cmd.Immediate(1))
			param2 := v.getParam(v.Ip+2, cmd.Immediate(2))
			address := v.getParam(v.Ip+3, true)
			if param1 < param2 {
				v.Code[address] = 1
			} else {
				v.Code[address] = 0
			}
			v.Ip += 4

		case 8:
			// equal
			param1 := v.getParam(v.Ip+1, cmd.Immediate(1))
			param2 := v.getParam(v.Ip+2, cmd.Immediate(2))
			address := v.getParam(v.Ip+3, true)
			if param1 == param2 {
				v.Code[address] = 1
			} else {
				v.Code[address] = 0
			}
			v.Ip += 4

		case 99:
			// halt
			return

		default:
			log.Fatalf("not an opcode %v", cmd)
		}
	}
}

func (v *VM) getParam(address int, immediate bool) int {
	param := v.Code[address]
	if immediate {
		return param
	}
	return v.Code[param]
}
