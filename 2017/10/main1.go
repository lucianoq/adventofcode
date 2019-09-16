package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	buf, _ := ioutil.ReadAll(os.Stdin)
	input := parseInts(buf)

	fmt.Println(hash(input))
}

func parseInts(buf []byte) []uint8 {
	list := make([]uint8, 0)
	items := strings.Split(string(buf), ",")
	for _, i := range items {
		n, _ := strconv.ParseUint(i, 10, 8)
		list = append(list, uint8(n))
	}
	return list
}

func hash(input []uint8) int {
	list := createList(size)

	pos := 0
	skipSize := 0

	for _, l := range input {
		list = knot(list, pos, int(l))
		pos += int(l) + skipSize
		skipSize++
	}

	return int(list[0]) * int(list[1])
}
