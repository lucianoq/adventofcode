package main

import "fmt"

func main() {
	facts, rules := parseInput()
	z := GetNumber("z", facts, rules)
	fmt.Println(z)
}
