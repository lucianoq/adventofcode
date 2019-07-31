package main

import (
	"fmt"
)

const (
	input      = "1321131112"
	iterations = 40
)

func main() {
	line := input
	for i := 0; i < iterations; i++ {
		line = EncodeIterBuf(line)
		//line = EncodeIter(line)
		//line = EncodeRecursive(line)
	}
	fmt.Println(len(line))
}
