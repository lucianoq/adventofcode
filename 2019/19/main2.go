package main

import "fmt"

func main() {
	// skip the beginning (too small beam)
	// and go in a upper left corner that for sure
	// is left of the stream
	y := 20
	x := 0

	for {
		// Upper Left corner is left of the beam
		// move right
		if !Beam(x, y) {
			x++
			continue
		}

		// Upper right corner is out (right) of the beam
		// move down
		if !Beam(x+99, y) {
			y++
			continue
		}

		// Lower left corner is out (left) of the beam
		// move right
		if !Beam(x, y+99) {
			x++
			continue
		}

		// If the three above are satisfied,
		// Lower right corner is always in the beam

		fmt.Println(x*10000 + y)
		return
	}
}

func Beam(x, y int) bool {
	vm := NewVM("input")

	go vm.Run()

	vm.Input <- x
	vm.Input <- y

	beam := <-vm.Output

	return beam == 1
}
