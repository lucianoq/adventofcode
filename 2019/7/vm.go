package main

import (
	"io/ioutil"
	"log"
	"math"
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
	Code   []int // source code and memory, list of int
	Ip     int   // Instruction Pointer
	Input  <-chan int
	output chan<- int
}

func (v *VM) Load(filename string) {
	buf, err := ioutil.ReadFile(filename)
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
			v.Code[address] = <-v.Input
			v.Ip += 2

		case 4:
			// write
			param1 := v.getParam(v.Ip+1, cmd.Immediate(1))
			v.output <- param1
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

func NewVM(filename string, input <-chan int, output chan<- int) *VM {
	vm := &VM{
		Code:   nil,
		Ip:     0,
		Input:  input,
		output: output,
	}

	vm.Load(filename)

	return vm
}
