package main

import "fmt"

func main() {
	m := parse()
	m.North()
	fmt.Println(m.Load())
}
