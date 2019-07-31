package main

import "fmt"

func main() {
	program := parse()

	reg := map[string]uint64{"a": 1, "b": 0}

	fmt.Println(run(program, reg))
}
