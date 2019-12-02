package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse() []int {
	buf, _ := ioutil.ReadAll(os.Stdin)
	listStr := strings.Split(strings.TrimSpace(string(buf)), ",")
	listInt := make([]int, len(listStr))

	for i := range listStr {
		listInt[i], _ = strconv.Atoi(listStr[i])
	}
	return listInt
}

func run(code []int) int {
	pos := 0
	for {
		switch code[pos] {
		case 99:
			return code[0]
		case 1:
			code[code[pos+3]] = code[code[pos+1]] + code[code[pos+2]]
			pos += 4
		case 2:
			code[code[pos+3]] = code[code[pos+1]] * code[code[pos+2]]
			pos += 4
		default:
			log.Fatal("wrong opcode")
		}
	}
}
