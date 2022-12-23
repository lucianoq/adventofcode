package main

import (
	"fmt"
)

func main() {
	parse()

	for i := 0; ; i++ {
		if !run() {
			fmt.Println(i + 1)
			return
		}
	}
}
