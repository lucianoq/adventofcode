package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var pat = []int{0, 1, 0, -1}

func pattern(i, j int) int {
	return pat[(j+1)/(i+1)%4]
}

func main() {
	list := parse()

	for phase := 0; phase < 100; phase++ {

		newList := make([]int, len(list))
		for i := range list {
			for j := range list {
				newList[i] += list[j] * pattern(i, j)
			}
			newList[i] = abs(newList[i]) % 10
		}
		list = newList
	}

	str := ""
	for _, i := range list[:8] {
		str += strconv.Itoa(i)
	}
	fmt.Println(str)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parse() []int {
	buf, _ := ioutil.ReadAll(os.Stdin)
	buf = bytes.TrimSpace(buf)
	ints := make([]int, len(buf))
	for i, b := range buf {
		ints[i] = int(b - '0')
	}
	return ints
}
