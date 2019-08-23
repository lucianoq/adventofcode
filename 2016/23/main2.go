package main

import "fmt"

const eggs = 12

func main() {

	// factorial of eggs
	x := 1
	for i := 1; i <= eggs; i++ {
		x *= i
	}

	// 86 and 78 are constants found in line 20 and 21 of the input
	x += 86 * 78

	fmt.Println(x)

}

// Normal approach required ~ 59 seconds
func main2() {
	parse()

	reg[0] = eggs // put eggs in a

	for cursor = 0; cursor < len(cmds); cursor++ {
		cmds[cursor].Run()
	}

	fmt.Println(reg[0])
}
