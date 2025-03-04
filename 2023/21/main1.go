package main

import "fmt"

func main() {
	m := parse()

	fmt.Println(Reach(m, 64))
}
