package main

import (
	"bytes"
	"io/ioutil"
	"os"
)

var pat = [4]int8{0, 1, 0, -1}

func pattern(i, j int) int8 {
	if j < i {
		return 0
	}
	if j >= i && j < 2*i+1 {
		return 1
	}
	return pat[(j+1)/(i+1)%4]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parse() []int8 {
	buf, _ := ioutil.ReadAll(os.Stdin)
	buf = bytes.TrimSpace(buf)
	ints := make([]int8, len(buf))
	for i, b := range buf {
		ints[i] = int8(b - '0')
	}
	return ints
}

func run100phases(list []int8) []int8 {
	size := len(list)
	for phase := 0; phase < 100; phase++ {
		newList := make([]int8, size)
		for i := 0; i < size; i++ {
			var sum int
			for j := i; j < size; j++ {
				sum += int(pattern(i, j) * list[j])
			}

			newList[i] = int8(abs(sum) % 10)
		}
		list = newList
	}
	return list
}
