package main

import (
	"bytes"
	"io/ioutil"
	"math"
	"os"
)

func parse() []int {
	buf, _ := ioutil.ReadAll(os.Stdin)
	buf = bytes.TrimSpace(buf)
	ints := make([]int, len(buf))
	for i, b := range buf {
		ints[i] = int(b - '0')
	}
	return ints
}

func toInt(list []int) int {
	num := 0
	for i, v := range list {
		num += int(math.Pow10(len(list)-i-1)) * v
	}
	return num
}
