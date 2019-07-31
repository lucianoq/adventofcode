package main

import (
	"fmt"
)

const (
	input  = "yzbqklnj"
	prefix = "00000"
)

func main() {
	inputs := make(chan string, 100)
	output := make(chan int)

	go generator(inputs)
	for j := 0; j < 10; j++ {
		go worker(inputs, output, prefix)
	}

	i := <-output
	fmt.Println(i)
}
