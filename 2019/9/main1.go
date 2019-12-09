package main

import "fmt"

func main() {

	input := make(chan int, 0)
	output := make(chan int, 0)

	go func() {
		NewVM("input", input, output).Run()
		close(output)
	}()

	input <- 1

	for x := range output {
		fmt.Println(x)
	}
}
