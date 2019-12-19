package main

import "fmt"

func main() {
	sum := 0
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if Beam(x, y) {
				fmt.Print("#")
				sum++
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println(sum)
}

func Beam(x, y int) bool {
	vm := NewVM("input")

	go vm.Run()

	vm.Input <- x
	vm.Input <- y

	beam := <-vm.Output

	return beam == 1
}
