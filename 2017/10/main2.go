package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	buf, _ := ioutil.ReadAll(os.Stdin)
	inputStr := strings.TrimSpace(string(buf))
	fmt.Println(hash2(inputStr))
}

func hash2(inputStr string) string {
	list := createList(size)

	lengths := []uint8(inputStr)
	lengths = append(lengths, []uint8{17, 31, 73, 47, 23}...)

	var pos, skipSize int
	for i := 0; i < 64; i++ {
		for _, l := range lengths {
			list = knot(list, pos, int(l))

			pos += int(l) + skipSize
			skipSize++
		}
	}

	return toString(dense(list))
}

func dense(list []uint8) [block]uint8 {
	var res [block]uint8

	for i := 0; i < block; i++ {
		var acc uint8
		for j := 0; j < block; j++ {
			acc ^= list[i*block+j]
		}
		res[i] = acc
	}

	return res
}

func toString(list [block]uint8) string {
	s := ""
	for _, i := range list {
		s += fmt.Sprintf("%02x", i)
	}
	return s
}
