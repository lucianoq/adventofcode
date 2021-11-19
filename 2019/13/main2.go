package main

import (
	"fmt"
	"time"
)

const (
	Noop  = 0
	Left  = -1
	Right = 1
)

func main() {
	done := make(chan struct{})

	input := make(chan int, 0)
	output := make(chan int, 0)
	vm := NewVM("input", input, output)

	go func() {
		// play for free!
		vm.Code[0] = 2

		vm.Run()
		close(output)
	}()

	// output
	grid := make(map[C]Tile)

	var ballPos, paddlePos C

	// load grid
	for i := 0; i < 21*36; i++ {
		x, y, tile := <-output, <-output, Tile(<-output)
		grid[C{x, y}] = tile
		if tile == Ball {
			ballPos = C{x, y}
		}
		if tile == Paddle {
			paddlePos = C{x, y}
		}
	}

	go func() {

		var maxScore int

		// read changes
		for {
			x, open := <-output
			if !open {
				fmt.Println(maxScore)
				done <- struct{}{}
				return
			}
			y := <-output

			if x == -1 && y == 0 {
				score := <-output
				if score > maxScore {
					maxScore = score
				}
				continue
			}

			tile := Tile(<-output)
			grid[C{x, y}] = tile

			if tile == Ball {
				ballPos = C{x, y}
			}
			if tile == Paddle {
				paddlePos = C{x, y}
			}
		}
	}()

	// input
	go func() {
		for {
			switch {
			case ballPos.x < paddlePos.x:
				input <- Left
			case ballPos.x > paddlePos.x:
				input <- Right
			case ballPos.x == paddlePos.x:
				input <- Noop
			}

			time.Sleep(time.Millisecond)
		}
	}()

	<-done
}
