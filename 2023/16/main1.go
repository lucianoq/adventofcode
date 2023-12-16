package main

import "fmt"

func main() {
	m := parse()

	fmt.Println(energy(m, Beam{0, -1, E}))
}
