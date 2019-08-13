package main

import "fmt"

func main() {
	reg["c"] = 1
	run()
	fmt.Println(reg["a"])
}
