package main

import (
	"fmt"
)

const (
	input  = "yzbqklnj"
	prefix = "000000"
)

func main() {
	inputs := make(chan string, 1000)
	output := make(chan int)

	go generator(inputs)
	for j := 0; j < 30; j++ {
		go worker(inputs, output, prefix)
	}

	i := <-output
	fmt.Println(i)
}
