package main

import "fmt"

func main() {
	code := parse()
	code[1], code[2] = 12, 2
	fmt.Println(run(code))
}
