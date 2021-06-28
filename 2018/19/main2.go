package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	instructions := read(os.Stdin)

	register = [6]int{1, 0, 0, 0, 0, 0}
	ip = 0
	for {
		if ip < 0 || ip >= len(instructions) {
			log.Println("Halt")
			break
		}

		instructions[ip].Exec()
	}

	fmt.Println(register[0])
}
