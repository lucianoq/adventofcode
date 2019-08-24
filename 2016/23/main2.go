package main

import "fmt"

const eggs = 12

func main() {

	// 86 and 78 are constants found in line 20 and 21 of the input
	res := factorial(eggs) + 86*78

	fmt.Println(res)
}

// Normal approach requires ~59 seconds
func main2() {
	parse()

	reg[0] = eggs // put eggs in a

	for cursor = 0; cursor < len(cmds); cursor++ {
		cmds[cursor].Run()
	}

	fmt.Println(reg[0])
}

func factorial(n int) int {
	x := 1
	for i := 1; i <= n; i++ {
		x *= i
	}
	return x
}
