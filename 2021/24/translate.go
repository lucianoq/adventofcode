package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("package main")

	fmt.Println("func alu() bool {")
	fmt.Println("w := 0")
	fmt.Println("x := 0")
	fmt.Println("y := 0")
	fmt.Println("z := 0")

	scanner := bufio.NewScanner(os.Stdin)

	digit := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		ff := strings.Split(line, " ")

		switch ff[0] {
		case "inp":
			fmt.Printf("%s = input[%d]\n", ff[1], digit)
			digit++

		case "add":
			fmt.Printf("%s += %s\n", ff[1], ff[2])

		case "mul":
			if ff[2] == "0" {
				fmt.Printf("%s = 0\n", ff[1])
				continue
			}
			fmt.Printf("%s *= %s\n", ff[1], ff[2])

		case "div":
			if ff[2] == "1" {
				continue
			}
			fmt.Printf("%s /= %s\n", ff[1], ff[2])

		case "mod":
			fmt.Printf("%s %%= %s\n", ff[1], ff[2])

		case "eql":
			fmt.Printf("if %s == %s { %s = 1 } else { %s = 0 } \n", ff[1], ff[2], ff[1], ff[1])
		}
	}

	fmt.Println("}")
}
