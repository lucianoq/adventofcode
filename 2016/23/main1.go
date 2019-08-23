package main

import "fmt"

const eggs = 7

func main() {
	parse()

	reg[0] = eggs

	for cursor = 0; cursor < len(cmds); cursor++ {
		cmds[cursor].Run()
	}

	fmt.Println(reg[0])
}
