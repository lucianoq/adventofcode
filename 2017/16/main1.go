package main

import "fmt"

func main() {
	cmds := parse()
	s := []byte(input)
	for _, c := range cmds {
		c.Apply(s)
	}
	fmt.Println(string(s))
}
