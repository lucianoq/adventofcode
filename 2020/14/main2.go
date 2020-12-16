package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	memory := map[uint64]uint64{}

	scanner := bufio.NewScanner(os.Stdin)

	var mask string

	for scanner.Scan() {
		line := scanner.Text()

		ff := strings.Split(line, " = ")
		if ff[0] == "mask" {
			mask = ff[1]
			continue
		}

		address, _ := strconv.ParseUint(strings.Trim(ff[0], "me[]"), 10, 64)
		val, _ := strconv.ParseUint(ff[1], 10, 64)
		addresses := maskAddress(mask, address)

		for _, add := range addresses {
			memory[add] = val
		}
	}

	var count uint64
	for _, v := range memory {
		count += v
	}
	fmt.Println(count)
}

func maskAddress(mask string, address uint64) []uint64 {

	binAdd := []byte(fmt.Sprintf("%036b", address))

	for i := range mask {
		switch mask[i] {
		case 'X':
			binAdd[i] = 'X'
		case '1':
			binAdd[i] = '1'
		}
	}

	// expand results
	todo := make([]string, 0)
	done := make([]string, 0)

	todo = append(todo, string(binAdd))
	var curr string
	for len(todo) > 0 {
		curr, todo = todo[0], todo[1:]

		idx := strings.Index(curr, "X")

		if idx == -1 {
			done = append(done, curr)
			continue
		}

		todo = append(todo, curr[:idx]+"0"+curr[idx+1:])
		todo = append(todo, curr[:idx]+"1"+curr[idx+1:])
	}

	// convert results
	res := make([]uint64, 0, len(done))
	for _, s := range done {
		val, _ := strconv.ParseUint(s, 2, 64)
		res = append(res, val)
	}

	return res
}
