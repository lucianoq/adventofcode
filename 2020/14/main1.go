package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	memory := map[int]uint64{}

	scanner := bufio.NewScanner(os.Stdin)

	var mask0, mask1 uint64

	for scanner.Scan() {
		line := scanner.Text()

		ff := strings.Split(line, " = ")
		if ff[0] == "mask" {
			mask0, mask1 = parseMask(ff[1])
			continue
		}

		address, _ := strconv.Atoi(strings.Trim(ff[0], "me[]"))
		val, _ := strconv.ParseUint(ff[1], 10, 64)
		val &= mask0
		val |= mask1
		memory[address] = val
	}

	var count uint64
	for _, v := range memory {
		count += v
	}
	fmt.Println(count)
}

func parseMask(mask string) (uint64, uint64) {
	var to0, to1 uint64

	for i := 0; i < len(mask); i++ {
		switch mask[i] {
		case '1':
			to1 |= 1 << (35 - i)
		case 'X':
			to0 |= 1 << (35 - i)
		}
	}

	return to0, to1
}
